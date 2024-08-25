package main

import "fmt"

const STEP = 20.0
const LOWER = 0.0
const UPPER = 300.0

func main() {
	for fahr := LOWER; fahr <= UPPER; fahr += STEP {
		fmt.Printf("%3.0f %6.1f\n", fahr, (5.0/9.0)*(fahr-32.0))
	}
}
