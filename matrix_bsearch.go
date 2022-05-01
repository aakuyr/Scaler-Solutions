package main

import "fmt"

func findMinMax(A [][]int, R, C int) (int, int) {
	max := 0
	min := 1000000001

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if min > A[i][j] {
				min = A[i][j]
			}

			if max < A[i][j] {
				max = A[i][j]
			}
		}
	}

	return min, max
}

func lessMoreThanEle(A [][]int, R, C, B int) (int, int) {
	less := 0
	more := 0

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if B < A[i][j] {
				less++
			} else if B > A[i][j] {
				more++
			}
		}
	}

	return less, more
}

func findMedian(A [][]int) int {
	R := len(A)
	C := len(A[0])

	start, end := findMinMax(A, R, C)
	l := R * C
	fmt.Println(l / 2)
	for start <= end {
		mid := start + (end-start)/2
		more, less := lessMoreThanEle(A, R, C, mid)
		if more+less < l-1 {
			if more >= l/2+1 {
				start = mid + 1
			} else if less >= l/2+1 {
				end = mid - 1
			} else {
				return mid
			}
		} else {
			if less > more {
				end = mid - 1
			} else if less < more {
				start = mid + 1
			} else if less == more {
				return mid
			}
		}
	}

	return -1
}
func main() {
	A := [][]int{
		{1, 2, 2, 2, 2, 3, 3, 3, 3, 3},
	}
	fmt.Println(findMedian(A))
}
