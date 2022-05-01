package main

import "fmt"

func oldPhoneCode(A string, codes map[byte]string, combo string) []string {
	if len(A) == 0 {
		return []string{combo}
	}
	result := make([]string, 0)

	if A == "0" {
		if combo != "" {
			return []string{combo + "0"}
		} else {
			return []string{"0"}
		}
	} else if A == "1" {
		if combo != "" {
			return []string{combo + "1"}
		} else {
			return []string{"1"}
		}
	} else {
		for _, char := range codes[A[0]] {
			var subCombos []string
			if combo == "" {
				subCombos = oldPhoneCode(A[1:], codes, string(char))
			} else {
				subCombos = oldPhoneCode(A[1:], codes, combo+string(char))
			}
			for _, j := range subCombos {
				result = append(result, j)
			}

		}
	}

	return result

}

func main() {
	arr := "0123456789"
	codes := map[byte]string{
		arr[0]: "0",
		arr[1]: "1",
		arr[2]: "abc",
		arr[3]: "def",
		arr[4]: "ghi",
		arr[5]: "jkl",
		arr[6]: "mno",
		arr[7]: "pqrs",
		arr[8]: "tuv",
		arr[9]: "wxyz",
	}
	A := "111"
	fmt.Println(oldPhoneCode(A, codes, ""))

}
