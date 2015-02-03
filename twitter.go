package main

import (
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	MIN_REPLY_DELAY_MS = 120000 // Should be 2 minutes
	MAX_REPLY_DELAY_MS = 300000 // Should be 5 minutes
)

type Reply struct {
	tweetText       string
	senderHandle    string
	originalTweetId string
}

type Favorite struct {
	senderHandle    string
	originalTweetId int64
}

type TwitterConnector struct {
	api             *anaconda.TwitterApi
	outgoingFaves   chan Favorite
	outgoingReplies chan Reply
}

func NewTwitterConnector() (error, *TwitterConnector) {
	var (
		consumerSecret    string
		consumerKey       string
		accessToken       string
		accessTokenSecret string
	)
	// Get the environment variables for twitter api access
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	// If the environment variables are missing, we need to furnish them
	if consumerSecret == "" {
		return errors.New("TWITTER_CONSUMER_SECRET environment variable was missing"), nil
	}
	if consumerKey == "" {
		return errors.New("TWITTER_CONSUMER_KEY environment variable was missing"), nil
	}
	if accessToken == "" {
		return errors.New("TWITTER_ACCESS_TOKEN_SECRET environment variable was missing"), nil
	}
	if accessTokenSecret == "" {
		return errors.New("TWITTER_ACCESS_TOKEN environment variable was missing"), nil
	}
	// Create a client
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	tc := TwitterConnector{
		api:             anaconda.NewTwitterApi(accessToken, accessTokenSecret),
		outgoingReplies: make(chan Reply),
		outgoingFaves:   make(chan Favorite),
	}
	// Return the connector
	return nil, &tc
}

func (tc *TwitterConnector) listenForTweets() {
	// Create parameters for the request
	params := url.Values{}
	params.Set("track", "owlhacks,#owlhacks2015,hackathon,hackathons,hackru,hackcwru,hophacks")
	// Get dat stream
	stream, err := tc.api.PublicStreamFilter(params)
	if err != nil {
		log.Println("Could not start the twitter listener:\n\t", err)
		return
	} else {
		for {
			select {
			case incomingStreamItem, ok := <-stream.C:
				if ok {
					go tc.handleIncomingTweet(incomingStreamItem)
				} else {
					log.Fatalln("The tweet stream closed suddenly")
				}
			}
		}
	}
}

func (tc *TwitterConnector) sendReplies() {
	for {
		select {
		case reply, ok := <-tc.outgoingReplies:
			if ok {
				params := url.Values{}
				params.Set("in_reply_to_status_id", reply.originalTweetId)
				_, err := tc.api.PostTweet(reply.tweetText, params)

				if err == nil {
					log.Println("Replied to @" + reply.senderHandle + "'s tweet")
				} else {
					if strings.Contains(err.Error(), "Status is over 140") {
						log.Println("Could not reply to @" + reply.senderHandle + "'s tweet, because the reply was too long.")
					} else {
						log.Println("Could not reply to @" + reply.senderHandle + "'s tweet:\n->\t" + err.Error())
					}
				}
			} else {
				log.Println("The outgoing replies buffer closed")
				return
			}
		}
	}
}

func (tc *TwitterConnector) sendFavorites() {
	for {
		select {
		case fave, ok := <-tc.outgoingFaves:
			if ok {
				tc.api.Favorite(fave.originalTweetId)
				log.Println("Tweet from @" + fave.senderHandle + " favorited")
			} else {
				log.Println("The outgoing favorites buffer closed")
				return
			}
		}
	}
}

func (tc *TwitterConnector) hang() int64 {
	// Returns time waited
	waitTime := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(MAX_REPLY_DELAY_MS) + MIN_REPLY_DELAY_MS
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	return waitTime
}

func (tc *TwitterConnector) handleIncomingTweet(potentialTweet interface{}) {
	hasWaited := false
	tweet, ok := potentialTweet.(anaconda.Tweet)
	if ok {
		fromHandle := tweet.User.ScreenName
		isNotRetweet := (tweet.RetweetedStatus == nil)
		if fromHandle != "owlhacks" && isNotRetweet {
			tweetId := tweet.Id
			tweetText := tweet.Text
			responseTweet, isPunified := punify(tweetText)

			log.Println("Incoming tweet from @" + fromHandle)
			// Favorite the tweet if its about Owlhacks
			if strings.Contains(tweetText, "owlhacks") || strings.Contains(tweetText, "Owlhacks") {
				if !hasWaited {
					// Prepare a random wait time, to create some drama
					log.Println("Waiting to respond to @" + fromHandle + "...")
					waitTime := tc.hang()
					log.Println("Waited about " + strconv.FormatInt((waitTime/60000), 10) + " minutes to respond to @" + fromHandle)
					hasWaited = true
				}

				tc.outgoingFaves <- Favorite{
					senderHandle:    fromHandle,
					originalTweetId: tweetId,
				}
			}
			// Only if we could figure out a punified version of the tweet, continue
			if isPunified {
				tweetText := "@" + fromHandle + " \"" + responseTweet + "\" #owled"
				if len(tweetText) <= 140 {
					if !hasWaited {
						// Prepare a random wait time, to create some drama
						log.Println("Waiting to respond to @" + fromHandle + "...")
						waitTime := tc.hang()
						log.Println("Waited about " + strconv.FormatInt((waitTime/60000), 10) + " minutes to respond to @" + fromHandle)
						hasWaited = true
					}

					tweetId := tweet.Id
					// Send the response
					tc.outgoingReplies <- Reply{
						tweetText:       tweetText,
						senderHandle:    fromHandle,
						originalTweetId: strconv.FormatInt(tweetId, 10),
					}
				} else {
					log.Println("Our response to @" + fromHandle + "'s tweet was too long (" + strconv.Itoa(len(tweetText)) + ")")
				}
			}
			// If we haven't waited yet, the tweet was ignored
			if !hasWaited {
				log.Println("@" + fromHandle + "'s tweet was ignored")
			}
		}
	} else {
		log.Println("Could not read the tweet")
	}
}
