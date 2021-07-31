package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	fmt.Println(fact(num))
}

// fact = n * (n-1) * (n-2) * ... * 1

func fact(num int) int {
	if num == 0 {
		return 1
	}

	result := num

	for iter := num - 1; iter > 1; iter-- {
		result *= iter
	}

	return result
}
