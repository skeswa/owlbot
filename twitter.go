package main

import (
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"os"
)

type TwitterConnector struct {
	api TwitterApi
}

func NewTwitterConnector() (error, TwitterConnector) {
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
	return nil, tc
}

func (tc TwitterConnector) findOwlableTweets() {
}
