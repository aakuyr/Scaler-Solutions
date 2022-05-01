package main

import "fmt"

func excelColumn(A string, columnCodes map[byte]int) int {
	n := len(A)
	result := 0

	for i := 0; i < n; i++ {
		result *= 26
		result += columnCodes[A[i]]
	}

	return result
}

func createColCodes() map[byte]int {
	colCodes := make(map[byte]int)
	for i := 65; i < 65+26; i++ {
		colCodes[byte(i)] = i - 64
	}

	return colCodes
}

func main() {
	A := "AB"
	fmt.Println(excelColumn(A, createColCodes()))
}
