package main

import (
	"RAWSH/src/internal"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("$ ")
	for scanner.Scan() {
		line := scanner.Text()
		tokens, err := internal.Tokenize(line)

        if err != nil {
            fmt.Println(err)
        }

		for _, token := range tokens {
			fmt.Println(token)
		}

		fmt.Printf("$ ")
	}
}

// func handle_command(line string) {
//     if line == "" {
//         return
//     }
//
// 	pieces := strings.Split(strings.Trim(line, " "), " ")
//
//     cmd := exec.Command(pieces[0], pieces[1:]...)
// 	out, err := cmd.Output()
//
// 	if err != nil {
// 		fmt.Println("Err:", err)
// 	}
//
// 	fmt.Println(string(out))
// }
