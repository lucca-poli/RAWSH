package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type shellSession struct {
    aliases map[string]string
}

func CreateShellSession() shellSession {
    return shellSession {
        aliases: nil,
    }
}

// n sei oq fazer com isso
func getConfigPath() string {
    defaultPath := "~/.rawrc"

    configFile, err := os.Create(defaultPath)
    if err != nil {
        fmt.Println("Err:", err)
    }
    defer configFile.Close()

    return defaultPath
}

// e nem isso
func mapAliases(reader io.Reader) (map[string]string, error) {
    aliases := make(map[string]string)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        tokens := strings.SplitN(scanner.Text(), " ", 2)
        if tokens != nil && tokens[0] == "alias" {
            matched, err := regexp.MatchString(`^[a-z_]+=".*"$`, tokens[1])

            if err != nil {
                return aliases, err
            } else if !matched {
                return aliases, errors.New("Bad pattern")
            } else {
                alias := strings.SplitN(tokens[1], "=", 2)
                aliases[alias[0]] = alias[1][1:len(alias[1])-1] // Passing the string from inside the quotes to the alias
            }
        }
    }

    return aliases, nil
}
