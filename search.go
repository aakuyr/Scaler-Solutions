package main

import "fmt"

func bSearch(A []int, B int) int {
	start := 0
	end := len(A) - 1

	result := -1

	for start <= end {
		mid := start + (end-start)/2
		if A[mid] == B {
			return mid
		} else if A[mid] > B {
			end = mid - 1
		} else {
			result = mid
			start = mid + 1
		}
	}

	return result + 1
}

func main() {
	arr := []int{1, 3, 5, 6}
	fmt.Println(bSearch(arr, 5))
}
