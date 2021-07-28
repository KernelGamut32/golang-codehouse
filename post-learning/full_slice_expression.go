package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	output(x, y, z)
	y = append(y, 30, 40, 50)
	// output(x, y, z)
	x = append(x, 60)
	// output(x, y, z)
	z = append(z, 70)
	// output(x, y, z)
	z = append(z, 80)
	// output(x, y, z)
	z[0] = 3
	// output(x, y, z)
}

func output(x []int, y []int, z []int) {
	fmt.Println("x: ", x, " len: ", len(x), " cap: ", cap(x))
	fmt.Println("y: ", y, " len: ", len(y), " cap: ", cap(y))
	fmt.Println("z: ", z, " len: ", len(z), " cap: ", cap(z))
}
