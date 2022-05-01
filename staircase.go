package main

import "fmt"

func noOfBricks(n int) int {
	return n * (n + 1) / 2
}

func stairCase(n int) int {

	if n <= 0 {
		return 0
	}

	start := 1
	end := n - 1
	key := 1

	for start <= end {
		mid := start + (end-start)/2
		if noOfBricks(mid) < n {
			key = mid
			start = mid + 1
		} else if noOfBricks(mid) == n {
			key = mid
			break
		} else {
			end = mid - 1
		}
	}

	return key
}

func main() {
	fmt.Println(stairCase(0))
}
