package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Preparing the scanner object to get multiline
	scanner := bufio.NewScanner(os.Stdin)
	var message []string
	var nl, nw, nc = 0, 0, 0

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		nl++
		message = append(message, line)
	}

	err := scanner.Err()
	if err != nil {
		fmt.Println("Error reading standard input", err)
		os.Exit(0)
	}

	fmt.Println(message)

	for _, word := range message {
		nc += len(word)
		nw += len(strings.Split(word, " "))
	}
	fmt.Printf("Line count: %d\nWord Count: %d\nCharacter Count: %d", nl, nw, nc)

}
