package main

import (
	"regexp"
	"strings"
)

func FixText(text string) string {
	// Corriger les apostrophes entourées d'espaces.
	text = regexp.MustCompile(
		`'\s*(.*?)\s*'`,
	).ReplaceAllStringFunc(text, func(match string) string {
		inside := strings.TrimSpace(match[1 : len(match)-1])
		return "'" + inside + "'"
	})

	// Coller les groupes de ponctuation.
	text = regexp.MustCompile(
		`\s+([.,!?;:]+)`,
	).ReplaceAllString(text, "$1")

	// Ajouter un espace après la ponctuation si nécessaire.
	text = regexp.MustCompile(
		`([.,!?;:]+)([A-Za-z])`,
	).ReplaceAllString(text, "$1 $2")

	// Corriger les articles a et A.
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		word := strings.ToLower(
			strings.Trim(words[i+1], `"'([{`),
		)

		if words[i] == "a" || words[i] == "A" {
			if len(word) > 0 &&
				strings.ContainsRune("aeiouh", rune(word[0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}