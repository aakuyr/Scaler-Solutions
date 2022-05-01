package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseString(A string) string {
	n := len(A)
	B := make([]byte, n)
	for i := 0; i < n/2+n%2; i++ {
		B[i], B[n-i-1] = A[n-1-i], A[i]
	}
	return string(B)
}

func reverseWords(C string) string {
	C = strings.TrimSpace(C)
	regexp := regexp.MustCompile("\\s+")
	A := string(reverseString(regexp.ReplaceAllString(C, " ")))
	n := len(A)
	B := ""
	lastSpace := 0

	for i := 0; i < n; i++ {
		if A[i] == 32 {
			B += reverseString(A[lastSpace:i]) + " "
			lastSpace = i + 1
		} else if i == n-1 {
			B += reverseString(A[lastSpace : i+1])
		}
	}

	return B
}

func main() {
	A := "qxkpvo  f   w vdg t wqxy ln mbqmtwwbaegx   mskgtlenfnipsl bddjk znhksoewu zwh bd fqecoskmo"
	B := "fqecoskmo bd zwh znhksoewu bddjk mskgtlenfnipsl mbqmtwwbaegx ln wqxy t vdg w f qxkpvo"
	fmt.Println(reverseWords(A) == B)
	fmt.Println(reverseString("iob"))
}
