package main

import "fmt"

func findNextPosEle(arr []int, curr int) int {
	n := len(arr)

	for i := curr + 1; i < n; i++ {
		if arr[i] >= 0 {
			return i
		}
	}
	return n
}

func max(arr []int) int {
	n := len(arr)
	maxEle := arr[0]
	for i := 1; i < n; i++ {
		if maxEle < arr[i] {
			maxEle = arr[i]
		}
	}
	return maxEle
}

func findMaxSubArr(arr []int) int {
	n := len(arr)
	left := findNextPosEle(arr, -1)

	if left == n {
		return max(arr)
	}

	right := left

	sum := arr[left]
	total := sum

	for right < n {
		fmt.Println(left, right, sum, total)
		if sum < total {
			sum = total
		}

		for right < n-1 && total >= 0 {
			right++
			total += arr[right]

			if sum < total {
				sum = total
			}
		}

		if right == n-1 {
			return sum
		}

		if total < 0 {
			// fmt.Println(left, right, sum, total)
			left = findNextPosEle(arr, left+1)

			if left == n {
				return sum
			}

			right = left
			total = arr[left]
		}
	}

	return sum
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(arr)
	fmt.Println(findMaxSubArr(arr))
}
