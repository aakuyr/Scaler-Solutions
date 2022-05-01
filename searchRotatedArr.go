package main

import "fmt"

func findPivot(A []int) int {
	n := len(A)
	start := 0
	end := n - 1

	for start <= end {
		if A[start] < A[end] || start == end {
			break
		}
		mid := start + (end-start)/2
		if A[mid] < A[start] {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func bSearch(A []int, B int) int {
	n := len(A)
	start := 0
	end := n - 1
	result := -1

	if n == 1 && A[0] != B {
		return -1
	}

	for start <= end {
		mid := start + (end-start)/2

		if A[mid] == B {
			return mid
		} else if A[mid] > B {
			end = mid - 1
		} else if A[mid] < B {
			start = mid + 1
		}
	}
	return result
}

func findElementInRotatedArray(A []int, B int) int {
	n := len(A)
	pivot := findPivot(A)

	if pivot == 0 {

	}

	if A[pivot] == B {
		return pivot
	} else if B < A[pivot] {
		return -1
	}
	if B < A[n-1] {
		index := bSearch(A[pivot:n], B)
		if index == -1 {
			return -1
		} else {
			return pivot + index
		}
	} else if B > A[0] {
		return bSearch(A[0:pivot+1], B)
	}

	return -1
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := 3
	fmt.Println(findElementInRotatedArray(A, B))
}
