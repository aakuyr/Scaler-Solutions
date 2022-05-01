package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func trap(array []int) int {
	n := len(array)

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = array[0]
	rightMax[n-1] = array[n-1]

	for i := int(1); i < n; i++ {
		leftMax[i] = max(leftMax[i-1], array[i])
	}

	for i := int(n - 2); i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], array[i])
	}

	amount := int(0)

	for i := int(1); i < n-1; i++ {
		temp := min(leftMax[i-1], rightMax[i+1]) - array[i]
		if temp > 0 {
			amount += temp
		}
	}
	fmt.Println(amount)
	return amount
}

func main() {
	arr := []int{1, 2, 3, 0, 3, 1, 5}
	trap(arr)
}
