package main

import "fmt"

func transform(A []int) []int {
	n := len(A)
	pivot := 0
	k := 0

	for i := 0; i < n; i++ {
		if A[i] < pivot {
			A[i], A[k] = A[k], A[i]
			k++
		}
	}

	fmt.Println(k, A)
	if k == n || k == 0 {
		return A
	}

	for i := 1; i <= k && k < n; i += 2 {
		A[i], A[k] = A[k], A[i]
		k++
	}

	return A
}

func main() {
	A := []int{-1, 2, 3, -4, 5}
	fmt.Println(transform(A))
}
