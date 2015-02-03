package main

import (
	"log"
)

func main() {
	err, connector := NewTwitterConnector()
	if err != nil {
		log.Fatalln("Could not start the listener", err)
	} else {
		log.Println("Starting the twitter listener...\n")
		// Start listening for outgoing faves and tweets
		go connector.sendFavorites()
		go connector.sendReplies()
		// Open the twitter stream
		connector.listenForTweets()
		log.Println("Now exiting...\n")
	}
}
