package tokenizer

import (
	"fmt"
	"strings"
)

func Tokenize(line string) []string {
	sanitized_line := strings.Trim(line, " ")

	var tokens []string
	var token string
	var quoted *byte
	for i := 0; i < len(sanitized_line); i++ {
		token = fmt.Sprint(token, string(sanitized_line[i]))
		fmt.Println("token is:", token)
		if quoted != nil {
			fmt.Println("quoted is: ", string(*quoted))
		}

		if (sanitized_line[i] == '"' || sanitized_line[i] == '\'') && sanitized_line[i-1] != '\\' {
			if quoted != nil && sanitized_line[i] == *quoted {
				tokens = append(tokens, token)
				token = ""
				quoted = nil
			} else if quoted == nil {
				temp := sanitized_line[i]
				quoted = &temp
			}
		}
		if quoted != nil {
			continue
		}
		if sanitized_line[i] == ' ' || (i == len(sanitized_line)-1 && len(token) != 0) {
			tokens = append(tokens, token)
			token = ""
		}
	}

	return tokens
}
