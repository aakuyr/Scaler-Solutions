package main

import "fmt"

const mod = 1000000007

func power(A, B, M int) int {
	if B == 0 {
		return 1
	} else if B == 1 {
		return A % M
	}

	if B%2 == 0 {
		return power((A*A)%M, B/2, M) % M
	} else {
		return A * (power((A*A)%M, B/2, M) % M) % M
	}
}

func fatorialMod(B, M int) int {
	result := 1
	for i := 1; i <= B; i++ {
		fmt.Println(result, result < M)
		result = (result * i) % M
	}

	return result % M
}

func largePower(A, B int) int {
	pow := fatorialMod(B, mod-1)
	A %= mod
	return power(A, pow, mod)
}

func main() {
	fmt.Println(largePower(2, 27))
}
