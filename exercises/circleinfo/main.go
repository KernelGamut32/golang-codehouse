package main

import (
	"fmt"
	"math"
)

func main() {
	area, circ := circleinfo(3.5)
	fmt.Printf("area = %.2f, circumference = %.2f\n", area, circ)
}

func circleinfo(r float64) (float64, float64) {
	return math.Pi * r * r, math.Pi * r * 2
}
