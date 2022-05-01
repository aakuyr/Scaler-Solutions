package main

import (
	"fmt"
	"sort"
)

func minAbsDiff(A []int) int {
	n := len(A)
	sort.Ints(A)
	minDiff := A[1] - A[0]

	for i := 0; i+1 < n; i++ {
		temp := A[i+1] - A[i]
		if minDiff > temp {
			minDiff = temp
		}
	}

	return minDiff
}

func main() {
	A := []int{10, 20, 30, 4, 5}
	fmt.Println(minAbsDiff(A))
}
