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

		switch {
		case word == "(hex)" || word == "(bin)":
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

		case isCaseCommand(word):
			if len(result) > 0 {
				count := 1

				if i+1 < len(words) &&
					strings.HasPrefix(words[i+1], "(") == false &&
					isNumber(words[i+1]) {
					count, _ = strconv.Atoi(words[i+1])
					i++
				}

				if count > len(result) {
					count = len(result)
				}

				start := len(result) - count

				for j := start; j < len(result); j++ {
					result[j] = changeCase(
						result[j], word,
					)
				}
			}

		default:
			result = append(result, word)
		}
	}

	return FixText(strings.Join(result, " "))
}

func isCaseCommand(word string) bool {
	return word == "(up)" ||
		word == "(low)" ||
		word == "(cap)" ||
		word == "(up,)" ||
		word == "(low,)" ||
		word == "(cap,)"
}

func isNumber(word string) bool {
	_, err := strconv.Atoi(word)
	return err == nil
}

func changeCase(word, command string) string {
	switch {
	case strings.HasPrefix(command, "(up"):
		return strings.ToUpper(word)

	case strings.HasPrefix(command, "(low"):
		return strings.ToLower(word)

	case strings.HasPrefix(command, "(cap"):
		runes := []rune(strings.ToLower(word))

		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}

		return string(runes)
	}

	return word
}