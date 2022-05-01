package main

import "fmt"

func secondMax(A []int) int {
	n := len(A)
	max := A[0]
	count := 0

	for i := 0; i < n; i++ {
		if max < A[i] {
			max = A[i]
		} else if max == A[i] {
			count++
		}
	}

	if count == n {
		return 0
	}

	secondMax := A[0]

	for i := 0; i < n; i++ {
		if secondMax < A[i] && A[i] < max {
			secondMax = A[i]
		}
	}

	return secondMax
}

func main() {
	A := []int{5, 5, 5, 5, 5}
	fmt.Println(secondMax(A))
}
