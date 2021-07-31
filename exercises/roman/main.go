package main

import (
	"fmt"
	"strings"
)

var translations = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func main() {
	var number string

	fmt.Print("Enter a Roman Number: ")
	fmt.Scanln(&number)
	fmt.Printf("%s is equivalent %d\n", number, translate(strings.ToUpper(number)))
}

// number = "MMLXV"
func translate(number string) int {
	arabicVals := make([]int, len(number)+1)
	for index, digit := range number {
		if value, present := translations[digit]; present {
			arabicVals[index] = value
		} else {
			fmt.Printf("bad digit: %c\n", digit)
		}
	}
	fmt.Println("Before:", arabicVals)
	total := 0
	for index := 0; index < len(number); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}
	fmt.Println("After:", arabicVals)
	return total
}
