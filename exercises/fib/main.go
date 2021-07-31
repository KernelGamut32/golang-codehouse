package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int

	fmt.Print("How many Fibonacci numbers? ")
	fmt.Scanln(&num)
	if result, err := fib(num); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// See https://en.wikipedia.org/wiki/Fibonacci_number

func fib(num int) ([]int, error) {
	if num < 2 {
		return nil, errors.New("main: invalid number of digits requested")
	}
	nums := make([]int, num)
	nums[0], nums[1] = 1, 1

	for i := 2; i < num; i++ {
		nums[i] = nums[i-2] + nums[i-1]
	}
	return nums, nil
}
