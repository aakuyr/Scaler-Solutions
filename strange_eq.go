package main

import (
	"fmt"
	"math"
)

func findNoOfBits(n int) float64 {
	bits := float64(1)
	for n > 1 {
		n >>= 1
		bits++
	}

	return bits
}

func solve(a int) int {
	x := int(math.Pow(2, findNoOfBits(a)) - 1)
	y := int(math.Pow(2, findNoOfBits(a)))

	fmt.Println(x, y)

	x = x ^ a

	fmt.Println(x, y)

	return x ^ y
}

func main() {
	fmt.Println(solve(1))
}
