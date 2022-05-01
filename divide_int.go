package main

import "fmt"

func divide(A, B int) int {
	result := 0
	var C, D int

	if A > 0 {
		C = A
	} else {
		C = -1 * A
	}

	if B > 0 {
		D = B
	} else {
		D = -1 * B
	}

	for C >= D {
		temp := uint(0)
		for uint(D)<<(temp+1) <= uint(C) {
			temp++
		}
		C -= D << temp
		result += 1 << temp
	}

	if A < 0 && B < 0 {
		return result
	} else if A < 0 || B < 0 {
		return result * -1
	}
	return result
}

func main() {
	fmt.Println(divide(-2147483648, -1))
}
