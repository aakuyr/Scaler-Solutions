package main

import "fmt"

func smallestPrefix(A, B string) string {
	n := len(A)
	m := len(B)

	i := 1
	j := 0
	prefix := []byte{A[0]}
	for j < m && i < n && j < 1 {
		if A[i] < B[j] {
			prefix = append(prefix, A[i])
			i++
		} else {
			prefix = append(prefix, B[j])
			j++
		}
	}

	if i == 0 {
		prefix = append(prefix, A[i])
	} else if j == 0 {
		prefix = append(prefix, B[j])
	}

	return string(prefix)
}

func main() {
	A := "tom"
	B := "riddle"
	fmt.Println(smallestPrefix(A, B))
}
