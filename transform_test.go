package main

import (
	"strconv"
	"strings"
	"unicode"
)

func ProcessText(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(hex)" || word == "(bin)" {
			if len(result) > 0 {
				index := len(result) - 1
				base := 16

				if word == "(bin)" {
					base = 2
				}

				number, err := strconv.ParseInt(
					result[index], base, 64,
				)

				if err == nil {
					result[index] = strconv.FormatInt(number, 10)
				}
			}
			continue
		}

		command := ""
		count := 1

		switch word {
		case "(up)", "(up,":
			command = "up"
		case "(low)", "(low,":
			command = "low"
		case "(cap)", "(cap,":
			command = "cap"
		}

		if command != "" {
			if strings.HasSuffix(word, ",") &&
				i+1 < len(words) {
				n := strings.TrimSuffix(words[i+1], ")")
				if value, err := strconv.Atoi(n); err == nil {
					count = value
					i++
				}
			}

			if count > len(result) {
				count = len(result)
			}

			start := len(result) - count

			for j := start; j < len(result); j++ {
				result[j] = changeCase(result[j], command)
			}

			continue
		}

		result = append(result, word)
	}

	return FixText(strings.Join(result, " "))
}

func changeCase(word, command string) string {
	switch command {
	case "up":
		return strings.ToUpper(word)

	case "low":
		return strings.ToLower(word)

	case "cap":
		runes := []rune(strings.ToLower(word))

		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}

		return string(runes)
	}

	return word
}
