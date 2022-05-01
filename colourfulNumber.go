package main

import (
	"fmt"
	"strconv"
)

func getIntFromString(A string) int {
	if integer, err := strconv.Atoi(A); err == nil {
		return integer
	} else {
		panic("Error converting string to in")
	}
}

func colourfulNo(A int) int {
	digits := strconv.Itoa(A)
	n := len(digits)
	m := make(map[int]string)

	if n == 1 {
		return 1
	}

	for i := 0; i < n; i++ {
		ele := string(digits[i])
		product := getIntFromString(ele)
		if product <= 1 {
			return 0
		}
		if _, ok := m[product]; ok == true {
			return 0
		}
		m[product] = ele
	}

	for i := 2; i <= n-1; i++ {
		start := 0
		end := start + i - 1
		product := 1
		for k := start; k < end; k++ {
			ele := getIntFromString(string(digits[k]))
			product *= ele
		}

		for end < n {
			startEle := getIntFromString(string(digits[start]))
			endEle := getIntFromString(string(digits[end]))
			product = product * endEle

			if _, ok := m[product]; ok == true {
				return 0
			}
			m[product] = digits[start : end+1]
			product /= startEle
			start++
			end++

		}
	}
	return 1
}

func main() {
	fmt.Println(colourfulNo(345))
}
