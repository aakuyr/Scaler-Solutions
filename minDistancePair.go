package main

import "fmt"

func min(A, B int) int {
	if A < B {
		return A
	} else {
		return B
	}
}

func minPairs(A []int) int {
	n := len(A)
	m := make(map[int]int)
	minDistance := n

	for i := 0; i < n; i++ {
		_, ok := m[A[i]]
		if ok == false {
			m[A[i]] = i
		} else {
			minDistance = min(minDistance, i-m[A[i]])
			m[A[i]] = i
		}
	}

	if minDistance < n {
		return minDistance
	} else {
		return -1
	}
}

func main() {
	A := []int{7, 1, 3, 4, 1, 7}
	fmt.Println(minPairs(A))
}
