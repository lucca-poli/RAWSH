package tokenizer

import (
	"errors"
	"fmt"
	"strings"
)

func Tokenize(line string) ([]string, error) {
	spaceParsed := strings.SplitAfter(strings.Trim(line, " "), " ")
	var operatorParsed []string

	for _, parsed := range spaceParsed {
		operatorParsed = append(operatorParsed, parseOperators(parsed)...)
	}

	fullParsed, err := aggregateQuotes(operatorParsed)

	var tokens []string
	for _, token := range fullParsed {
		token = strings.Trim(token, " ")

		if token != "" {
			tokens = append(tokens, token)
		}
	}

	return tokens, err
}

func aggregateQuotes(parses []string) ([]string, error) {
	var tokens []string
	var quote rune
	var token, unquotedStr string

	for _, parsed := range parses {
		quote, unquotedStr = unquoteToken(quote, parsed)
		token = fmt.Sprint(token, unquotedStr)

		if quote == 0 {
			tokens = append(tokens, token)
			token = ""
		}
	}

	if quote != 0 {
		return tokens, errors.New("Bad quoting")
	}

	return tokens, nil
}

func unquoteToken(lastQuote rune, currentToken string) (rune, string) {
	var unquotedToken string

	for i, char := range currentToken {
		switch {
		case lastQuote != 0 && char == lastQuote:
			if char == '\'' {
				lastQuote = 0
			}
			if char == '"' && (i == 0 || currentToken[i-1] != '\\') {
				lastQuote = 0
			}
		case lastQuote == 0 && (char == '\'' || char == '"') && (i == 0 || currentToken[i-1] != '\\'):
			lastQuote = char
		default:
			unquotedToken = fmt.Sprint(unquotedToken, string(char))
		}
	}

	return lastQuote, unquotedToken
}

func parseOperators(s string) []string {
	var tokens []string
	var operator string
	lastOperatorIdx := -1

	for i, char := range s {

		if isMetachar(char) {
			if operator == "" {
				tokens = append(tokens, s[lastOperatorIdx+1:i])
			}

			operator = fmt.Sprint(operator, string(char))
		}

		if operator != "" && !isMetachar(char) {
			tokens = append(tokens, operator)
			operator = ""
			lastOperatorIdx = i - 1
		}
	}
	tokens = append(tokens, s[lastOperatorIdx+1:])

	return tokens
}

func isMetachar(char rune) bool {
	metachars := []rune{'|', '<', '>', '&', '(', ')', ';'}

	for _, c := range metachars {
		if char == c {
			return true
		}
	}

	return false
}

