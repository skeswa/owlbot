package main

import (
	"strings"
)

const (
	OWL_REPLACEMENT_LIMIT = 3
)

func punify(text string) (string, bool) {
	punifiedText := text
	replaceCount := 0
	// Replace "all"
	if strings.Contains(punifiedText, "all") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "all", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(punifiedText, "All") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "All", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(punifiedText, "al") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "al", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(punifiedText, "Al") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "Al", "OWL", 1)
		replaceCount += 1
	}
	if strings.Contains(punifiedText, "I'll") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "I'll", "OWL", 1)
		replaceCount += 1
	}
	// Trim the tweet
	punifiedText = strings.TrimSpace(punifiedText)
	// Check if there's an owl pun
	if strings.Contains(punifiedText, "OWL") {
		return punifiedText, true
	} else {
		return "", false
	}
}
