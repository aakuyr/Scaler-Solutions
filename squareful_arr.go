package main

import (
	"fmt"
	"math"
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

func isPerfectSquare(A int) bool {
	n := float64(A)
	s := math.Sqrt(n)
	if math.Ceil(s) == math.Floor(s) {
		return true
	}
	return false
}

func isSquarefulArray(A []int) bool {
	return true
}

func main() {
	fmt.Println(isPerfectSquare(4))
}
