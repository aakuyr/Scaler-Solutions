package main

import (
	"fmt"
	"strings"
)

func solve(A string) string {
	dateStrArr := strings.Split(A, " ")
	fmt.Println(dateStrArr)
	day := getDay(dateStrArr[0])
	month := getMonth(dateStrArr[1])
	year := getYear(dateStrArr[2])

	date := year + "-" + month + "-" + day

	return date
}

func getDay(A string) string {
	n := len(A)
	result := []byte{}
	for i := 0; i < n; i++ {
		if A[i] >= 48 && A[i] <= 57 {
			result = append(result, A[i])
		}
	}
	if len(result) == 1 {
		return "0" + string(result)
	} else {
		return string(result)
	}
}

func getMonth(A string) string {
	switch A {
	case "Jan":
		return "01"
	case "Feb":
		return "02"
	case "Mar":
		return "03"
	case "Apr":
		return "04"
	case "May":
		return "05"
	case "Jun":
		return "06"
	case "Jul":
		return "07"
	case "Aug":
		return "08"
	case "Sep":
		return "09"
	case "Oct":
		return "10"
	case "Nov":
		return "11"
	case "Dec":
		return "12"
	default:
		return "-1"
	}
}

func getYear(A string) string {
	return A
}

func main() {
	fmt.Println(solve("3rd July 3252"))
}
