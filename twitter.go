package main

import (
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type TwitterConnector struct {
	api *anaconda.TwitterApi
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
		api: anaconda.NewTwitterApi(accessToken, accessTokenSecret),
	}
	// Return the connector
	return nil, &tc
}

func (tc *TwitterConnector) listenForTweets() error {
	// Create parameters for the request
	params := url.Values{}
	params.Set("track", "#owlhacks,#owlhacks2015,hackathon,hackathons,#hackru")
	// Get dat stream
	stream, err := tc.api.PublicStreamFilter(params)
	if err != nil {
		return err
	} else {
		for {
			select {
			case incomingStreamItem := <-stream.C:
				tc.handleIncomingTweet(incomingStreamItem)
			}
		}
	}
}

func (tc *TwitterConnector) handleIncomingTweet(potentialTweet interface{}) {
	tweet, ok := potentialTweet.(anaconda.Tweet)
	if ok {
		fromHandle := tweet.User.ScreenName
		if fromHandle != "owlhacks" {
			tweetId := tweet.Id
			tweetText := tweet.Text
			responseTweet, isPunified := punify(tweetText)
			// Favorite the tweet if its about Owlhacks
			if strings.Contains(tweetText, "owlhacks") && strings.Contains(tweetText, "Owlhacks") {
				tc.api.Favorite(tweetId)
				log.Println("Favorited tweet from @" + fromHandle)
			}
			// Only if we could figure out a punified version of the tweet, continue
			if isPunified {
				tweetId := tweet.Id
				fromName := tweet.User.Name
				// Send the response
				params := url.Values{}
				params.Set("in_reply_to_status_id", strconv.FormatInt(tweetId, 10))
				_, err := tc.api.PostTweet("@"+fromHandle+" \""+responseTweet+"\" #owled", params)
				if err != nil {
					log.Println("Could not post a reply:", err)
				} else {
					log.Println("New tweet from " + fromName + ":\n-> \"" + tweetText + "\"\n<- \"" + responseTweet + "\"\n")
				}
			}
		}
	} else {
		log.Println("Could not read the tweet")
	}
}
