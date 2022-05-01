package main

import (
	"fmt"
)

func binSearch(A []int, num int) int {
	n := len(A)
	start := 0
	end := n - 1

	result := -1

	for start <= end {
		mid := start + (end-start)/2
		if A[mid] == num {
			result = mid
			break
		} else if A[mid] < num {
			result = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return A[result]
}

func main() {
	A := []int{2, 4, 6, 8, 10}
	fmt.Println(binSearch(A, 11))
}
