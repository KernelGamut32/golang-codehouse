package main

import (
	"fmt"
	"strconv"
)

func main() {
	testNums := [...]int{3, 4, 5, 15}
	for index, num := range testNums {
		fmt.Printf("%d: %s\n", index, fizzbuzz2(num))
		// fmt.Println(fizzbuzz(num))
	}
}

// func fizzbuzz(num int) string {
// 	if num%15 == 0 {
// 		return "fizzbuzz"
// 	}

// 	if num%3 == 0 {
// 		return "fizz"
// 	}

// 	if num%5 == 0 {
// 		return "buzz"
// 	}

// 	return strconv.Itoa(num)
// }

func fizzbuzz2(num int) string {
	switch {
	case num%15 == 0:
		return "fizzbuzz"
	case num%3 == 0:
		return "fizz"
	case num%5 == 0:
		return "buzz"
	default:
		return strconv.Itoa(num)
	}
}
