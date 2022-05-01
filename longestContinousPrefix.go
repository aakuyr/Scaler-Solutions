package main

import "fmt"

func longestCommonPrefix(A []string) string {
	n := len(A)
	result := []byte{}
	for p := 0; p < len(A[0]); p++ {
		prefixChar := A[0][p]
		for i := 1; i < n; i++ {
			if p >= len(A[i]) {
				p = len(A[0])
				break
			} else if prefixChar != A[i][p] {
				p = len(A[0])
				break
			}
		}
		if p < len(A[0]) {
			result = append(result, prefixChar)
		}
	}

	return string(result)
}

func main() {
	A := []string{"abcd", "aze	"}
	fmt.Println(longestCommonPrefix(A))
}
