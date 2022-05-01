package main

import "fmt"

func maxone(A []int, B int) []int {
	l := 0
	r := 0

	n := len(A)

	maxLength := 0
	gLeft := 0
	gRight := 0
	for r < n && l < n {
		if A[r] == 0 {
			B--
		}

		if r-l > maxLength && B >= 0 {
			maxLength = r - l
			gLeft = l
			gRight = r
		}

		if B < 0 {
			for A[l] == 1 {
				l++
			}
			l++
			B++
		}
		r++
	}

	result := make([]int, gRight-gLeft+1)
	for i := 0; i <= gRight-gLeft; i++ {
		result[i] = gLeft + i
	}

	return result
}

func main() {
	X := []int{0, 1, 1, 1}
	fmt.Println(maxone(X, 0))
}
