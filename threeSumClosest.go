package main

import (
	"fmt"
	"math"
	"sort"
)

func matches(B, result, tempResult int) int {
	existingAns := float64(B - result)
	newAns := float64(B - tempResult)
	if math.Abs(existingAns) > math.Abs(newAns) {
		return tempResult
	} else {
		return result
	}

	return result
}

func twoClosestMatches(A []int, B, start, end int) int {
	i := start
	j := end

	result := A[start] + A[end]

	for i < j {
		tempResult := A[i] + A[j]

		if tempResult < B {
			i++
		} else if tempResult > B {
			j--
		} else {
			return B
		}

		result = matches(B, result, tempResult)
	}

	return result
}

func threeSumClosest(A []int, B int) int {
	sort.Ints(A)
	n := len(A)

	fmt.Println(A)

	result := A[0] + A[1] + A[2]

	i := 0
	j := n - 1

	for i+1 < j {
		tempResult := twoClosestMatches(A, B-A[i], i+1, j)
		result = matches(B, result, A[i]+tempResult)
		i++
	}

	return result
}

func main() {
	A := []int{-4, -8, -10, -9, -1, 1, -2, 0, -8, -2}
	fmt.Println(threeSumClosest(A, 0))
}
