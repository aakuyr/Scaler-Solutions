package main

import "fmt"

func findNextPosEle(array []int, curr int) int {
	n := len(array)
	for i := curr; i < n; i++ {
		if array[i] >= 0 {
			return i
		}
	}
	return n
}

func findMaxSubArr(arr []int) []int {

	n := len(arr)
	left := findNextPosEle(arr, 0)

	if left == n {
		return nil
	}

	if left == n-1 {
		if arr[left] >= 0 {
			result := []int{arr[left]}
			return result
		}
	}

	right := left

	gleft := left
	gright := right

	sum := arr[left]
	temp := sum

	for right < n {
		if sum < temp {
			sum = temp
			gleft, gright = left, right
		}

		if sum == temp {
			if gright-gleft < right-left {
				gleft, gright = left, right
			}
		}

		if right >= n-1 {
			break
		}

		if arr[right+1] >= 0 {
			right++
			temp += arr[right]

		} else {
			left = findNextPosEle(arr, right+1)

			if left == n {
				break
			}

			right = left
			temp = arr[left]
		}
	}

	result := make([]int, gright-gleft+1)

	for i := 0; i <= gright-gleft; i++ {
		result[i] = arr[i+gleft]
	}

	return result
}

func main() {
	arr := []int{0, -1, 0, 0, 0}
	fmt.Println(findMaxSubArr(arr))
}
