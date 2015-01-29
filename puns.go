package main

import (
	"strings"
)

const (
	OWL_REPLACEMENT_LIMIT = 3
)

func punify(text string) string {
	punifiedText := text
	replaceCount := 0
	// Replace "all"
	if strings.Contains(text, "all") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "all", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(text, "All") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "All", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(text, "al") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "al", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(text, "Al") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "Al", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(text, "I'll") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "I'll", "OWL", 1)
		replaceCount += 1
	}
	// Get rid of the hashtag
	punifiedText = strings.Replace(punifiedText, "#owlhacks2015", "", -1)
	// Trim the tweet
	punifiedText = strings.TrimSpace(punifiedText)

	// Return that shite
	return punifiedText
}
