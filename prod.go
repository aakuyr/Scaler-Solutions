package main

import (
	"fmt"
)

func min(nums ...int) int {
	min := nums[0]
	for _, value := range nums {
		if min > value {
			min = value
		}
	}
	return min
}

func main() {
	fmt.Println(min(3, 2, 1))
}
