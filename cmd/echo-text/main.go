package main

import "fmt"

func main() {
	var input string

	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading std input", err)
	}

	fmt.Println(input)
}
