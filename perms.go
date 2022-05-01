package main

import (
	"fmt"
)

func inArray(A []int, B int) bool {
	for i := range A {
		if A[i] == B {
			return true
		}
	}
	return false
}

func perms(A []int, P []int, pos int) [][]int {
	n := len(A)
	R := [][]int{}

	if pos == n {
		temp := make([]int, len(P))
		copy(temp, P)
		R = append(R, temp)
		return R
	}

	for i := 0; i < n; i++ {
		if inArray(P, A[i]) {
			continue
		}
		P[pos] = A[i]
		R = append(R, perms(A, P, pos+1)...)
		P[pos] = 0
	}

	return R
}

func main() {
	A := []int{1, 2, 3, 4}
	l := len(A)
	fmt.Println(perms(A, make([]int, l), 0))
}
