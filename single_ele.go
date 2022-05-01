package main

import "fmt"

func singleEle(A []int) int {
	n := len(A)
	start := 0
	end := n - 1

	for start <= end {
		if start == end {
			return A[start]
		}

		mid := start + (end-start)/2

		if mid%2 == 0 {
			if A[mid] == A[mid+1] {
				start = mid + 2
			} else if A[mid] == A[mid-1] {
				end = mid - 2
			} else {
				return A[mid]
			}
		} else {
			if A[mid] == A[mid+1] {
				end = mid - 1
			} else if A[mid] == A[mid-1] {
				start = mid + 1
			} else {
				return A[mid]
			}
		}
	}

	return -1
}

func main() {
	A := []int{1, 1, 7}
	fmt.Println(singleEle(A))
}
