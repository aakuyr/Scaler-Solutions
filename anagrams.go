package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func hashString(A string, letterCodes map[byte]int) int {
	result := 0
	A = sortString(A)
	c := 1

	for i := 0; i < len(A); i++ {
		result += c * letterCodes[A[i]]
		c *= 31
	}

	return result
}

func anagrams(A []string, letterCodes map[byte]int) [][]int {
	n := len(A)
	hashOfStrings := make(map[int][]int)
	for i := 0; i < n; i++ {
		stringHash := hashString(A[i], letterCodes)
		if _, ok := hashOfStrings[stringHash]; ok == false {
			hashOfStrings[stringHash] = make([]int, 0)
			hashOfStrings[stringHash] = append(hashOfStrings[stringHash], i+1)
		} else {
			hashOfStrings[stringHash] = append(hashOfStrings[stringHash], i+1)
		}
	}

	result := [][]int{}
	index := 0
	for _, value := range hashOfStrings {
		result = append(result, []int{})
		for i := range value {
			result[index] = append(result[index], value[i])
		}
		index++
	}

	fmt.Println(result)
	return result
}

func createLetterCode() map[byte]int {
	A := "abcdefghijklmnopqrstuvwxyz"
	code := 1
	letterCodes := make(map[byte]int)
	for i := 0; i < 26; i++ {
		letterCodes[A[i]] = code
		code++
	}

	return letterCodes
}

func main() {
	A := []string{"cde", "bee"}
	letterCodes := createLetterCode()
	anagrams(A, letterCodes)
}
