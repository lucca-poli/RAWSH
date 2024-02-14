package tokenizer

import (
	"fmt"
	"strings"
)

func Tokenize(line string) ([]string, error) {
	sanitized_line := strings.Trim(line, " ")

	var tokens []string
	var token string
	var quoted *byte
	for i := 0; i < len(sanitized_line); i++ {
		token = fmt.Sprint(token, string(sanitized_line[i]))

		if quoted != nil || ((sanitized_line[i] == '"' || sanitized_line[i] == '\'') && (i == 0 || sanitized_line[i-1] != '\\')) {
			if quoted != nil && sanitized_line[i] == *quoted {
				token = token[:len(token)-1]
				tokens = append(tokens, token)
				token = ""
				quoted = nil
			} else if quoted == nil {
				token = ""
				temp := sanitized_line[i]
				quoted = &temp
			}
		} else if sanitized_line[i] == ' ' || (i == len(sanitized_line)-1 && len(token) != 0) {
			if strings.Trim(token, " ") == "" {
				token = ""
				continue
			}
			tokens = append(tokens, strings.Trim(token, " "))
			token = ""
		}
	}

	return tokens, nil
}
