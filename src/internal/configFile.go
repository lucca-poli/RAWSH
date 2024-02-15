package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func GetConfigPath() string {
    defaultPath := "~/.rawrc"

    configFile, err := os.Create(defaultPath)
    if err != nil {
        fmt.Println("Err:", err)
    }
    defer configFile.Close()

    return defaultPath
}

func MapAliases(reader io.Reader) (map[string]string, error) {
    aliases := make(map[string]string)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        tokens := strings.SplitN(scanner.Text(), " ", 2)
        if tokens != nil && tokens[0] == "alias" {
            matched, err := regexp.MatchString(`^[a-z_]+=".*"$`, tokens[1])

            if err != nil || !matched {
                return aliases, err
            } else {
                alias := strings.SplitN(tokens[1], "=", 2)
                aliases[alias[0]] = alias[1][1:len(alias[1])-1] // Passing the string from inside the quotes to the alias
            }
        }
    }

    return aliases, nil
}