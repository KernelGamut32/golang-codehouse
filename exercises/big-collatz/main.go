package main

import (
	"fmt"
	"math/big"
)

// Supports calcs against REALLY, REALLY, REALLY, REALLY, REALLY big numbers
// Compare to the "collatz" version

func main() {
	var num string
	fmt.Print("Enter number: ")
	fmt.Scanln(&num)
	collatz(num)
}

func collatz(num string) bool {
	number := new(big.Int)
	number, ok := number.SetString(num, 10)
	if !ok {
		return false
	}

	big_one := big.NewInt(1)
	big_two := big.NewInt(2)
	big_three := big.NewInt(3)
	if number.Cmp(big_one) < 0 {
		return false
	}
	var count = 1
	for ; number.Cmp(big_one) > 0; count++ {
		fmt.Printf("%d ", number)
		if number.Bit(0) == 0 {
			number = number.Div(number, big_two)
		} else {
			number = number.Mul(number, big_three).Add(number, big_one)
		}
		if count%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("1")
	fmt.Printf("Total steps: %d\n", count)
	return true
}
