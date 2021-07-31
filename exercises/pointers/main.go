package main

import "fmt"

func main() {
	str := "Golang"
	doubler(&str)
	fmt.Println(str)
}

func doubler(s *string) {
	fmt.Println(s)
	doubleHelper(s)
}

func doubleHelper(s *string) {
	*s += *s
}
