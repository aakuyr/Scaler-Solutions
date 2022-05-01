package main

import "fmt"

func peakEle(A []int) int {
	n := len(A)
	start := 0
	end := n - 1

	if n == 2 {
		if A[0] > A[1] {
			return A[0]
		} else {
			return A[1]
		}
	}

	for start <= end {
		if start == end {
			return A[start]
		}

		mid := start + (end-start)/2

		if mid == start {
			if A[start] > A[end] {
				return A[start]
			} else {
				return A[end]
			}
		}

		if A[mid] >= A[mid+1] && A[mid] >= A[mid-1] {
			return A[mid]
		} else if A[mid] >= A[mid+1] {
			end = mid - 1
		} else if A[mid] >= A[mid-1] {
			start = mid + 1
		}

	}

	return -1
}

func main() {
	A := []int{1, 5, 2, 1, 1}
	fmt.Println(peakEle(A))
}
