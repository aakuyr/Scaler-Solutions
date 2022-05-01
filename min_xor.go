package main

import (
	"fmt"
	"sort"
)

func min_xor(A []int) int {

	n := len(A)
	sort.Ints(A)
	xor := A[n-1] ^ 0
	temp := A[n-1]

	for i := n - 2; i >= 0; i-- {
		if xor > temp^A[i] {
			xor = temp ^ A[i]
		}
		temp = A[i]
	}

	return xor
}

func main() {

	arr := []int{0, 2, 5, 7}
	fmt.Println(min_xor(arr))

}
