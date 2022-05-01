package main

import "fmt"

func merge_sorted_arrays(A, B []int) []int {
	m := len(A)
	n := len(B)

	i := 0
	j := 0
	k := 0

	C := make([]int, m+n)
	for i < m && j < n {
		if A[i] <= B[j] {
			C[k] = A[i]
			i++
		} else {
			C[k] = B[j]
			j++
		}
		k++
	}

	if i < m {
		for i < m {
			C[k] = A[i]
			i++
			k++
		}
	} else if j < n {
		for j < n {
			C[k] = B[j]
			j++
			k++
		}
	}

	return C
}

func main() {
	A := []int{4, 7, 9}
	B := []int{1, 2, 3, 5, 6, 7, 10}

	fmt.Println(merge_sorted_arrays(A, B))
}
