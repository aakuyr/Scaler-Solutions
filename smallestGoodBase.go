package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func getIntString(A float64) string {
	return strconv.FormatFloat(A, 'f', -1, 64)
}

func getFloat64(A string) float64 {
	if number, err := strconv.ParseFloat(A, 64); err == nil {
		return number
	} else {
		panic("Error converting string to float")
	}
}
func smallestGoodBase(A string) string {
	number := getFloat64(A)
	maxBits := math.Floor(math.Log2(number))
	for i := maxBits; i >= float64(2); i-- {
		k := math.Pow(number, math.Pow(i, -1))
		k = math.Floor(k)
		numerator := math.Pow(k, i+1) - 1
		denominator := k - 1
		if number*denominator == numerator {
			return getIntString(k)
		}
	}
	if number-1 == number {
		res, ok := new(big.Int).SetString(A, 10)
		if ok == false {
			panic("Can't convert to bigint")
		}
		res = res.Sub(res, big.NewInt(1))
		return res.String()
	}
	number = number - 1
	return getIntString(number)
}

func main() {
	fmt.Printf(smallestGoodBase("40782091"))
}
