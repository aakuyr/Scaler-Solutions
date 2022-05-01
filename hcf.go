package main

import "fmt"

func hcf(A, B int) int {
	if B < A {
		A, B = B, A
	}
	for A > 0 {
		A, B = B%A, A
	}
	return B
}

func main() {
	fmt.Println(hcf(20, 9))
}
