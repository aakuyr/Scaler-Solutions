package main

import "fmt"

const (
	mod = 1000000007
)

func countRectangles(A []int, B int) int {
	n := len(A)
	i := 0
	j := n - 1

	count := 0

	for i <= j {
		if A[i]*A[j] >= B {
			j--
		} else {
			count += (2*(j-i) + 1) % mod
			i++
		}
	}

	return count
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	fmt.Println(countRectangles(A, 5))
}
