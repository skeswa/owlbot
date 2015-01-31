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
	if strings.Contains(punifiedText, "about") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "about", "OWLbout", 1)
		replaceCount += 1
	}
	if strings.Contains(punifiedText, "About") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "About", "OWLbout", 1)
		replaceCount += 1
	}
	if strings.Contains(punifiedText, "who") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "who", "hoot", 1)
		replaceCount += 1
	}
	if strings.Contains(punifiedText, "Who") && replaceCount < OWL_REPLACEMENT_LIMIT {
		punifiedText = strings.Replace(punifiedText, "Who", "Hoot", 1)
		replaceCount += 1
	}
	// Trim the tweet
	punifiedText = strings.TrimSpace(punifiedText)
	// Check if there's an owl pun
	if replaceCount >= OWL_REPLACEMENT_LIMIT {
		return punifiedText, true
	} else {
		return "", false
	}
}
