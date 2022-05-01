package main

import "fmt"

const modulus = 1000000007

func combination(A int) int64 {
	if A <= 1 {
		return 0
	} else if A == 2 {
		return 1
	}
	n := 2
	m := A - 2
	result := float64(1)

	for i := m + 1; i <= A; i++ {
		result *= float64(i)
	}

	result /= float64(n)

	return int64(result) % modulus
}

func mod(A []int, B int) int {
	modValues := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if _, ok := modValues[A[i]%B]; ok == false {
			modValues[A[i]%B] = 1
		} else {
			modValues[A[i]%B]++
		}
	}

	count := 0

	for i := 0; i <= B/2; i++ {
		if i == 0 || i == B-i {
			count += int(combination(modValues[i]))
		} else {
			count += (modValues[i] * modValues[B-i]) % modulus
		}
	}

	return count % modulus
}

func main() {
	fmt.Println(mod([]int{5, 17, 100, 11}, 28))
}
