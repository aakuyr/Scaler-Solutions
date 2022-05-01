package main

import (
	"fmt"
	"sort"
)

const (
	mod = 1000000007
)

func distinctPairs(A []int, B int) int {
	n := len(A)
	sort.Ints(A)

	i := 0
	j := 1

	count := 0

	for i < j && j < n {
		if A[j]-A[i] < B {
			j++
		} else if A[j]-A[i] > B {
			i++
		} else {
			count = (count + 1) % mod
			for j+1 < n && A[j] == A[j+1] {
				j++
			}
			for i < j && A[i] == A[i+1] {
				i++
			}
			if j+1 < n {
				i++
			} else {
				j++
			}
		}

		if i == j {
			j++
		}
	}
	return count
}

func main() {
	// A := []int{1, 8, 2, 8, 8, 8, 8, 4, 4, 6, 10, 10, 9, 2, 9, 3, 7}
	A := []int{1, 1, 1}
	fmt.Println(distinctPairs(A, 0))
}
