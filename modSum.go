package main

import "fmt"

/*
	Here for all i, 1 <= A[i] <= 1000 but length of array is 100000.
	Hence there will be duplicate elements in the array, and it is best to loop
	over the range from 1 to the max element of array to save iterations.
	There seems to be no way to get the answer in less than O(N**2) as we need
	A % M + B % M which can't be gotten by just summing the elements and applying modulus
*/
func max(A []int) int {
	n := len(A)
	max := 0

	for i := 0; i < n; i++ {
		if max < A[i] {
			max = A[i]
		}
	}

	return max
}

func modSum(A []int) int {
	n := len(A)
	m := max(A)
	eleFreq := make(map[int]int)

	for i := 0; i < n; i++ {
		if _, ok := eleFreq[A[i]]; ok == false {
			eleFreq[A[i]] = 1
		} else {
			eleFreq[A[i]]++
		}
	}

	ans := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= m; j++ {
			ans += eleFreq[i] * eleFreq[j] * (i % j)
		}
	}

	return ans
}

func main() {
	A := []int{1, 2, 3}
	fmt.Println(modSum(A))
}
