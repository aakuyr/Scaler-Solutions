package main

import "fmt"

func isPossible(A []int, B, min int) bool {
	n := len(A)
	sum := 0
	noOfStudents := 1

	for i := 0; i < n; i++ {

		if A[i] > min {
			return false
		}

		if sum+A[i] > min {
			noOfStudents++
			sum = A[i]

			if noOfStudents > B {
				return false
			}
		} else {
			sum += A[i]
		}
	}

	return true
}

func min(A, B int) int {
	if A < B {
		return A
	} else {
		return B
	}
}

func allocateMin(A []int, B int) int {
	n := len(A)
	sum := 0
	result := 100000000

	if n < B {
		return -1
	}

	for i := 0; i < n; i++ {
		sum += A[i]
	}

	for i := 0; i < n; i++ {
		if A[i] > sum {
			return -1
		}
	}

	start := 0
	end := sum

	for start <= end {
		mid := start + (end-start)/2
		fmt.Println(start, end, mid)
		if isPossible(A, B, mid) {
			result = min(result, mid)
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return result
}

func main() {
	A := []int{31, 14, 19, 75}
	fmt.Println(allocateMin(A, 12))
}
