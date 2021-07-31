package main

import "fmt"

func main() {
	// Use one of the following to see how bad input is handled
	// nums := []float64{}
	// var nums []float64
	nums := []float64{54.2, 67.9, 19.3, 23.4, -19.3, -54.2, -23.4}
	if min, max, avg, ok := minmaxavg(nums); ok {
		fmt.Printf("%4f, %4f, %.4f\n", min, max, avg)
	} else {
		fmt.Println("bad slice input")
	}
}

func minmaxavg(nums []float64) (float64, float64, float64, bool) {
	if len(nums) == 0 {
		return 0.0, 0.0, 0.0, false
	}

	min, max, total := nums[0], nums[0], 0.0
	for _, value := range nums {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
		total += value
	}
	return min, max, total / float64(len(nums)), true
}
