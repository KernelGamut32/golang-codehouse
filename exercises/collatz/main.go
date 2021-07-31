package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	collatz(num)
}

func collatz(num int) bool {
	if num < 1 {
		fmt.Println("Ouch!")
		return false
	}
	count := 1
	// if it's even, divide it by 2
	// if it's odd, multiply it by 3 and add 1
	for ; num > 1; count++ {
		fmt.Printf("%5d", num)
		if num%2 == 0 {
			num /= 2
			// num = num / 2
		} else {
			num = 3*num + 1
		}
		if count%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("    1")
	fmt.Printf("Total steps: %d\n", count)
	return true
}
