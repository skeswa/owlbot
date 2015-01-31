package main

import (
	"strings"
)

const (
	OWL_REPLACEMENT_LIMIT = 3
)

var (
	replacementMap = map[string]string{
		"all":   "OWL",
		"All":   "OWL",
		"al":    "OWL",
		"ol":    "OWL",
		"Al":    "OWL",
		"I'll":  "OWL",
		"about": "abHOOT",
		"About": "AbHOOT",
	}
)

func punify(text string) (string, bool) {
	punifiedText := text
	replaceCount := 0

	for what, with := range replacementMap {
		if strings.Contains(punifiedText, what) {
			punifiedText = strings.Replace(punifiedText, what, with, -1)
			replaceCount += 1

			if replaceCount > OWL_REPLACEMENT_LIMIT {
				return punifiedText, true
			}
		}
	}

	return punifiedText, (replaceCount > 0)
}
