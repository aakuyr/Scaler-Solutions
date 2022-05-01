package main

import "fmt"

func similarString(S string, N int) bool {
	n := len(S)
	m := make(map[byte]int)

	for i := 0; i < n; i++ {
		if _, ok := m[S[i]]; ok == false {
			m[S[i]] = 1
		} else {
			m[S[i]]++
		}
	}

	fmt.Println(m)

	for _, value := range m {
		if value%N != 0 {
			return false
		}
	}

	return true
}

func main() {
	A := "abcabcabc"
	fmt.Println(similarString(A, 3))
}
