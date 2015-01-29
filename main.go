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
		connector.listenForTweets()
	}
}
