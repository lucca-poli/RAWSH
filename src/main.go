package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("$ ")
	for scanner.Scan() {
		line := scanner.Text()
        handle_command(line)

		fmt.Printf("$ ")
	}
}

func handle_command(line string) {
    if line == "" {
        return 
    }

	pieces := strings.Split(strings.Trim(line, " "), " ")

    cmd := exec.Command(pieces[0], pieces[1:]...)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("Err:", err)
	}

	fmt.Println(string(out))
}
