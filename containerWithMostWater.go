package main

import "fmt"

func min(A, B int) int {
	if A < B {
		return A
	} else {
		return B
	}
}

func max(A, B int) int {
	if A > B {
		return A
	} else {
		return B
	}
}

func calcArea(A []int, i, j int) int {
	width := j - i
	height := min(A[i], A[j])
	return height * width
}

func findMaxArea(A []int) int {
	n := len(A)

	i := 0
	j := n - 1

	result := 0

	for i < j {
		result = max(result, calcArea(A, i, j))
		if A[i] < A[j] {
			i++
		} else if A[i] > A[j] {
			j--
		} else {
			i++
			j--
		}
	}

	return result
}

func main() {
	A := []int{1, 5, 4, 3}
	fmt.Println(findMaxArea(A))
}
