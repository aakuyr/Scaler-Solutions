package main

import "fmt"

func main() {
	a := "sai is best"
	for i := range a {
		fmt.Printf("%T\n", a[i])
	}
}
