package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ops = map[rune]func(int, int) int{
	'+': add,
	'-': sub,
	'*': mul,
	'/': div,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var op rune
	var num1, num2 int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.ToLower(line) == "quit now" {
			return
		}
		line = strings.Replace(line, " ", "", -1)
		if n, _ := fmt.Sscanf(line, "%d%c%d", &num1, &op, &num2); n == 3 {
			if f, present := ops[op]; present {
				fmt.Println(f(num1, num2))
			} else {
				fmt.Println("bad operator:", op)
			}
		} else {
			fmt.Println("bad line:", line)
		}
	}
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}
