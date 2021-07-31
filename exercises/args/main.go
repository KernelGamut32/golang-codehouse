package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) < 3 {
		fmt.Println("Format is <program> num1 num2")
		return
	}

	var num1 int
	var err error

	if num1, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println(err)
	}
	num2, _ := strconv.Atoi(os.Args[2])
	fmt.Println(num1 + num2)
}
