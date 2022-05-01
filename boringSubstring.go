package main

import (
	"fmt"
	"math"
)

func maxOddCharacter(A string) int {
	n := len(A)
	max := 0

	for i := 0; i < n; i++ {
		if A[i]%2 == 1 {
			if max < int(A[i]) {
				max = int(A[i])
			}
		}
	}

	return max
}

func minOddCharacter(A string) int {
	n := len(A)
	min := 200

	for i := 0; i < n; i++ {
		if A[i]%2 == 1 {
			if min > int(A[i]) {
				min = int(A[i])
			}
		}
	}

	return min
}

func minEvenCharacter(A string) int {
	n := len(A)
	min := 200

	for i := 0; i < n; i++ {
		if A[i]%2 == 0 {
			if min > int(A[i]) {
				min = int(A[i])
			}
		}
	}

	return min
}

func maxEvenCharacter(A string) int {
	n := len(A)
	max := 0

	for i := 0; i < n; i++ {
		if A[i]%2 == 0 {
			if max < int(A[i]) {
				max = int(A[i])
			}
		}
	}

	return max
}

func boringSubString(A string) bool {
	minEven := minEvenCharacter(A)
	maxEven := maxEvenCharacter(A)
	minOdd := minOddCharacter(A)
	maxOdd := maxOddCharacter(A)

	diff1 := math.Abs(float64(maxEven - minOdd))
	diff2 := math.Abs(float64(maxOdd - minEven))

	if diff1 == 1 && diff2 == 1 {
		return false
	}

	return true
}

func main() {
	result := boringSubString("nmnn")
	fmt.Println(result)
}
