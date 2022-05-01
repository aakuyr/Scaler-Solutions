package main

import "fmt"

func calcRightTriangles(A, B []int) int {
	countOfX := make(map[int]int)
	countOfY := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if _, ok := countOfX[A[i]]; ok == false {
			countOfX[A[i]] = 1
		} else {
			countOfX[A[i]]++
		}

		if _, ok := countOfY[B[i]]; ok == false {
			countOfY[B[i]] = 1
		} else {
			countOfY[B[i]]++
		}
	}

	count := 0
	for i := 0; i < len(A); i++ {
		if countOfX[A[i]] >= 1 && countOfY[B[i]] >= 1 {
			count += (countOfX[A[i]] - 1) * (countOfY[B[i]] - 1)
		}
	}

	return count
}
func main() {
	A := []int{1, 1, 2, 3, 3}
	B := []int{1, 2, 1, 2, 1}
	fmt.Println(calcRightTriangles(A, B))
}
