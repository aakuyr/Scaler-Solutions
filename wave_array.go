package main

import (
	"fmt"
	"sort"
)

func wave_array(A []int) []int {
	n := len(A)
	sort.Ints(A)

	for i := 0; i+1 < n; i += 2 {
		A[i], A[i+1] = A[i+1], A[i]
	}

	return A
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	fmt.Println(wave_array(A))
}
