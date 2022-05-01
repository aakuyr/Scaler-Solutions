package main

import "fmt"

func sqrt(A int) int {
	start := 0
	end := A
	result := 0

	for start <= end {
		mid := start + (end-start)/2
		if mid*mid == A {
			return mid
		} else if mid*mid < A {
			result = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return result
}

func main() {
	var A int
	fmt.Scanf("%d", &A)
	fmt.Println(sqrt(A))
}
