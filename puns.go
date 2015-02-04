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
		"Howl":  "OWL",
		"howl":  "OWL",
		"foul":  "OWL",
		"Foul":  "OWL",
		"owel":  "OWL",
		"about": "abHOOT",
		"About": "AbHOOT",
		"oot":   "HOOT",
	}
	// Black-listed puns
	exceptionMap = map[string]string{
		"reOWL":     "real",
		"ReOWL":     "Real",
		"tOWLk":     "talk",
		"TOWLk":     "Talk",
		"chOWLk":    "chalk",
		"ChOWLk":    "Chalk",
		"stOWLk":    "stalk",
		"StOWLk":    "Stalk",
		"wOWLk":     "walk",
		"WOWLk":     "Walk",
		"bOWLk":     "balk",
		"BOWLk":     "Balk",
		"disembOWL": "disembowel", //plsno
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
