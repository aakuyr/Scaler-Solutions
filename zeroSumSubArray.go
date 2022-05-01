package main

import "fmt"

func calcPrefixArray(A []int) []int {
	n := len(A)
	C := make([]int, n)
	sum := 0

	for i := 0; i < n; i++ {
		sum += A[i]
		C[i] = sum
	}

	return C
}

func max(A, B int) int {
	if A >= B {
		return A
	} else {
		return B
	}
}

func calcLongestZeroSumSubArray(A []int) int {
	n := len(A)
	P := calcPrefixArray(A)

	fmt.Println(P)

	m := make(map[int]int)

	length := 0

	for i := 0; i < n; i++ {
		if P[i] != 0 {
			_, ok := m[P[i]]
			if ok == false {
				m[P[i]] = i
			} else {
				fmt.Println(i, P[i], m[P[i]])
				length = max(i-m[P[i]]+1, length)
			}
		} else {
			length = max(length, i+1)
		}
	}

	if length > 1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	A := []int{-78, -97, -44, -18, -7, -26, 37, -76, -23, -35, 48, 9, 25, 62, -90, 27, -40, 18, 88, 82, 15, 96, 31, -2, -45, -48, 52, -78, -79, -76, -18, -88, -85, 58, -48, -48, -16, 77, -79, -89, -78, 27, 98, 53, -6, 43, 73, 38}
	fmt.Println(calcLongestZeroSumSubArray(A))
}
