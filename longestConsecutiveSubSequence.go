package main

import "fmt"

func exists(A map[int]bool, B int) bool {
	_, ok := A[B]
	return ok
}

func max(A, B int) int {
	if A > B {
		return A
	} else {
		return B
	}
}

func longestConsecutiveSubSeq(A []int) int {
	m := make(map[int]bool)
	for _, ele := range A {
		m[ele] = true
	}

	length := 0

	for _, ele := range A {
		if !exists(m, ele-1) {
			tempLength := 1
			for exists(m, ele+1) {
				tempLength++
				ele++
			}
			length = max(tempLength, length)
		}
	}

	return length
}

func main() {
	A := []int{1}
	fmt.Println(longestConsecutiveSubSeq(A))
}
