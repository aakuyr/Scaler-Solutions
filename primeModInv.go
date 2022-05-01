package main

import "fmt"

func gcd(A, B int) int {
	if A < B {
		A, B = B, A
	}
	for B > 0 {
		A, B = B, A%B
	}

	return A
}

func power(A, B, M int) int {
	fmt.Println(A, B, M)
	if B == 0 {
		return 1
	} else if B == 1 {
		return A % M
	}
	if B%2 == 0 {
		return power((A*A)%M, B/2, M) % M
	} else {
		return (A * power((A*A)%M, B/2, M) % M) % M
	}
}
func invMod(A, B int) int {
	gcd := gcd(A, B)
	if gcd == 1 {
		return power(A, B-2, B)
	} else {
		return -1
	}
}

func main() {
	fmt.Println(power(3734, 11, 13))
}
