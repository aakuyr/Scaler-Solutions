package main

import "fmt"

func findSubArray(A []int, B int) []int {
	n := len(A)

	i := 0
	j := 0

	sum := A[i]
	for i <= j && j < n {
		if sum == B {
			return A[i : j+1]
		} else if sum < B && j+1 < n {
			j++
			sum += A[j]
		} else {
			sum -= A[i]
			i++
		}
	}

	return []int{-1}
}

func main() {
	// A := []int{5, 10, 20, 100, 105}
	A := []int{1, 2, 3, 4, 5}
	fmt.Println(findSubArray(A, 5))
}
