package main

import (
	"fmt"
	"sort"
)

func findOddBit(arr []int) int {
	oddBit := -1
	count := 0

	for bit := 0; bit < 32; bit++ {
		count = 0
		for _, i := range arr {
			count += i >> bit & 1
		}
		if count%2 == 1 {
			oddBit = bit
			break
		}
	}
	return oddBit
}
func missingEle(arr []int) []int {
	n := len(arr)
	oddBit := findOddBit(arr)
	xor1 := 0
	xor2 := 0

	for i := 0; i < n; i++ {
		if temp := arr[i] >> oddBit; temp%2 == 0 {
			xor1 = xor1 ^ arr[i]
		} else {
			xor2 = xor2 ^ arr[i]
		}
	}
	result := []int{xor2, xor1}
	sort.Ints(result)
	return result
}

func main() {
	arr := []int{3, 3, 4, 2}
	fmt.Println(missingEle(arr))
}
