package main

import (
	"fmt"
	"math"
)

func getAbsValue(n int) int {
	return int(math.Abs(float64(n)))
}

func segregateNgatives(array []int) ([]int, int) {
	j := 0
	n := len(array)

	for i := 0; i < n; i++ {
		if array[i] <= 0 {
			array[i], array[j] = array[j], array[i]
			j++
		}
	}
	return array, j
}

func findMissingInt(array []int, j int) int {
	n := len(array)

	if j == n {
		return 1
	}

	max := array[j]
	answer := -1
	for i := j; i < n; i++ {
		ele := getAbsValue(array[i])
		if ele+j-1 < n {
			if array[ele+j-1] > 0 {
				array[ele+j-1] *= -1
			}
		}
	}
	fmt.Println(array)
	for i := j; i < n; i++ {
		if array[i] > 0 {
			answer = i - j + 1
			break
		}
		if ele := getAbsValue(array[i]); max < ele {
			max = ele
		}
	}

	if answer == -1 {
		return max + 1
	}

	return answer
}

func main() {
	arr := []int{-5}
	arr, j := segregateNgatives(arr)
	fmt.Println("The shift is " + fmt.Sprint(j))
	fmt.Println("The array is \n\n\n" + fmt.Sprint(arr) + "\n\n")
	fmt.Println(findMissingInt(arr, j))
}
