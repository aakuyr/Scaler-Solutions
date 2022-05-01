package main

func findNumberOfFactors(A, B int) int {
	count := 0
	divisor := B

	for divisor <= A {
		count += A / divisor
		divisor *= B
	}

	return count
}

func noOfZerosInFactorial(A int) int {
	numberOfTwos := findNumberOfFactors(A, 2)
	numberOfFives := findNumberOfFactors(A, 5)

	if numberOfTwos <= numberOfFives {
		return numberOfTwos
	}

	return numberOfFives
}
