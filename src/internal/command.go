package internal

import (
	"errors"
	"regexp"
	"strings"
)

type word = string
type operator = string

func Tokenizer(line string) ([]string, error) {
	var tokens []string
	trimmedLine := strings.Trim(line, " ")
	ignoredQuotesTokens := strings.SplitAfter(trimmedLine, " ")
	pattern := regexp.MustCompile(`('[^']*'|"[^"\\]*(?:\\.[^"\\]*)*")\s*`) // bash quoting rules
	matches := pattern.FindAllStringIndex(trimmedLine, -1)

	var splitCount int
	var quotedCount int
	for i := 0; i < len(trimmedLine); {
		if isQuoted(i, matches) {
			interval := matches[quotedCount]
			tokens = append(tokens, word(strings.Trim(trimmedLine[interval[0]:interval[1]], " ")))
			i = interval[1]
			for j := interval[0]; j < i; splitCount++ {
				j += len(ignoredQuotesTokens[splitCount])
			}
			if i != len(trimmedLine) && trimmedLine[i] != ' ' {
				splitCount--
			}
			quotedCount++
		} else if isOperator(trimmedLine[i : i+len(ignoredQuotesTokens[splitCount])]) {
			token := ignoredQuotesTokens[splitCount]
			tokens = append(tokens, operator(strings.Trim(token, " ")))
			i += len(token)
			splitCount++
		} else {
			token := ignoredQuotesTokens[splitCount]
			tokens = append(tokens, word(strings.Trim(token, " ")))
			i += len(token)
			splitCount++
		}

		if i == len(trimmedLine) { // update this to use regex instead
			lastWord := ignoredQuotesTokens[splitCount-1]
			for j, char := range lastWord {
				if (char == '\'' || char == '"') && (j == 0 || lastWord[j-1] != '\\') {
					return tokens, errors.New("Bad quoting pattern")
				}
			}
		}
	}

	return tokens, nil
}

func isQuoted(position int, intervals [][]int) bool {
	if intervals != nil {
		for _, match := range intervals {
			if position == match[0] {
				return true
			}
		}
	}
	return false
}

func isOperator(token string) bool {
	metacharacters := []rune{'|', '&', ';', '(', ')', '<', '>'}

	for i, char := range token {
		var isMeta bool
		for _, metaChar := range metacharacters {
			if char == metaChar {
				isMeta = true
			}
		}

		if isMeta && (i == 0 || token[i-1] != '\\') {
			return true
		}
	}

	return false
}
