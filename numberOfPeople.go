package main

import "math"

func findNumberOfPeopleLeft(A int) int {
	ans := math.Floor(math.Log2(float64(A)))
	return int(math.Pow(2, ans))
}
