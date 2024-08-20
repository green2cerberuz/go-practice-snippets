package main

import "fmt"

func main() {
	// print the fahrenheit sequence
	var fahr, celsius int

	lower, upper, step := 0, 300, 20

	fahr = lower

	for fahr <= upper {
		celsius = 5 * (fahr - 32) / 9
		fmt.Printf("%d\t%d\n", fahr, celsius)
		fahr += step
	}

}
