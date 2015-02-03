package main

import (
	"strings"
)

const (
	OWL_REPLACEMENT_LIMIT = 3
)

var (
	// Replacements that cause puns
	replacementMap = map[string]string{
		"all":   "OWL",
		"All":   "OWL",
		"al":    "OWL",
		"aul":   "OWL",
		"awl":   "OWL",
		"Al":    "OWL",
		"I'll":  "OWL",
		"about": "abHOOT",
		"About": "AbHOOT",
		"oot":   "HOOT",
	}
	// Black-listed puns
	exceptionMap = map[string]string{
		"reOWL": "real",
		"tOWLk": "talk",
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
				break
			}
		}
	}

	for what, with := range exceptionMap {
		if strings.Contains(punifiedText, what) {
			punifiedText = strings.Replace(punifiedText, what, with, -1)
			replaceCount -= 1

			if replaceCount <= 0 {
				break
			}
		}
	}

	return punifiedText, (replaceCount > 0)
}
