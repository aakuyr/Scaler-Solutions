package main

import "fmt"

func maxSubArrayLen(A []int, k int) int {
	n := len(A)
	left := 0
	right := 1
	sum := A[left]
	minLength := n

	if n == 2 {
		if A[0] > k || A[1] > k {
			return 0
		} else if A[0] == k || A[1] == k {
			return 1
		} else {
			return 2
		}
	}

	for right <= n {
		if right < n && sum <= k {
			sum += A[right]
			right++
		} else if left < right {
			fmt.Println(left, right, sum)
			if sum > k && minLength > right-left {
				minLength = right - left
			}
			sum -= A[left]
			left++
		} else if left == right {
			break
		}
	}

	if minLength == n {
		return n
	} else {
		return minLength - 1
	}
}

func main() {
	A := []int{1, 5}
	fmt.Println(maxSubArrayLen(A, 5))
}
