package main

func solve(A, B int) int {
	if B < A {
		A, B = B, A
	}
	return B - A
}
