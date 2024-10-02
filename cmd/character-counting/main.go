package main

import (
	"fmt"
	"os"
)

func main() {
	var input string

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Error reading standard input", err)
		os.Exit(0)
	}
	fmt.Println(len(input))
}
