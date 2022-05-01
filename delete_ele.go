package main

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	for b > 0 {
		a %= b
		a, b = b, a
	}

	return a
}

func gcdArr(arr []int) int {
	n := len(arr)

	result := gcd(arr[0], 0)

	for i := 1; i < n; i++ {
		result = gcd(result, arr[i])
	}

	return result
}
func coPrimeEle(arr []int) int {
	gcdOfArr := gcdArr(arr)

	if gcdOfArr == 1 {
		return 0
	}
	return -1

}
func main() {

}
