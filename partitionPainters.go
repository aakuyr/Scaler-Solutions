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

func partition(A []int, C, B int) int {
	n := len(A)
	sum := 0
	result := 100000000000

	for i := 0; i < n; i++ {
		A[i] = int(int64(A[i]) * int64(C) % int64(10000003))
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
		if isPossible(A, B, mid) {
			result = min(result, mid)
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return result % 10000003
}

func main() {
	A := []int{1000000, 1000000}
	fmt.Println(partition(A, 1000000, 1))
}
