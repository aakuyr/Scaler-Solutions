package main

import (
	"strconv"
)

func grayCode(A int) []string {
	if A == 1 {
		return []string{"0"}
	} else if A == 2 {
		return []string{"0", "1"}
	}

	grayCodes := grayCode(A - 1)
	result := []string{}
	n := len(grayCodes)

	for i := 0; i < n; i++ {
		result = append(result, "0"+grayCodes[i])
	}

	for i := n - 1; i >= 0; i-- {
		result = append(result, "1"+grayCodes[i])
	}

	return result
}

func returnInt(A []string) []int {
	n := len(A)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if integer, err := strconv.ParseInt(A[i], 2, 64); err == nil {
			arr[i] = int(integer)
		} else {
			panic("Invalid int")
		}
	}
	return arr
}
