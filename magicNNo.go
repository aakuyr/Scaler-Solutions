package main

import (
	"fmt"
	"math"
)

func calcMagicNumber(A int) int {
	result := 0

	for A > 0 {
		if A == 1 {
			result += 5
			break
		} else if A == 2 {
			result += 25
			break
		}
		count := int(math.Log2(float64(A)))
		result += int(math.Pow(5, float64(count+1)))
		A -= int(math.Pow(2, float64(count)))
	}

	return result
}

func main() {
	fmt.Println(calcMagicNumber(10))
}
