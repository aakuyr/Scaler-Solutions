package main

func firstNonRepeatingChar(A string) string {
	n := len(A)

	start := 0
	charFrequency := make(map[byte]int)

	B := make([]byte, n)

	for i := 0; i < n; i++ {
		if _, ok := charFrequency[A[i]]; ok == false {
			charFrequency[A[i]] = 1
		} else {
			charFrequency[A[i]]++
		}
		for charFrequency[A[start]] > 1 {
			start++
		}
		if start < i {
			B[i] = A[start]
		} else {
			B[i] = byte('#')
		}
	}

	return string(B)
}
