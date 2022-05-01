package main

import (
	"fmt"
)

func compare(A []int, B []int) bool {
	n := len(A)
	m := len(B)

	if n != m {
		return false
	}

	for i := 0; i < n; i++ {
		if A[i] != B[i] {
			return false
		}
	}

	return true
}

func appendComboToResult(R [][]int, P []int) [][]int {
	if P != nil {
		temp := make([]int, len(P))
		copy(temp, P)
		R = append(R, temp)
	}
	return R
}

func appendCombosToResult(R, C [][]int) [][]int {
	if C == nil {
		return R
	}

	for i := range C {
		flag := false

		for j := range R {
			if compare(C[i], R[j]) {
				flag = true
			}
		}

		if flag == false {
			R = appendComboToResult(R, C[i])
		}
	}

	return R
}

func comb(A []int, P []int, B int) [][]int {
	R := [][]int{}
	n := len(A)

	if B < 0 || n == 0 {
		return nil
	}

	if B == 0 {
		R = appendComboToResult(R, P)
		return R
	}

	newP := append(P, A[0])
	combination := comb(A, newP, B-A[0])
	R = appendCombosToResult(R, combination)

	combination = comb(A[1:], P, B)
	R = appendCombosToResult(R, combination)

	return R
}

func main() {
	A := []int{1, 2, 3}
	P := []int{}
	fmt.Println(comb(A, P, 5))
}
