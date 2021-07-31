package main

import (
	"fmt"
)

func main() {
	var num float64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	fmt.Println(sqrt(num))
}

func sqrt(num float64) float64 {
	currguess := 1.0

	for count := 1; count <= 10; count++ {
		// guess = guess − (guess² − x) / (2 * guess)
		currguess = currguess - (currguess*currguess-num)/(2*currguess)
		fmt.Println("Guess =", currguess)
	}

	return currguess
}
