package main

import "sort"

func consecutive_arr(A []int) int {
	sort.Ints(A)
	n := len(A)

	for i := 0; i+1 < n; i++ {
		if A[i]+1 != A[i+1] {

		}
	}
}
