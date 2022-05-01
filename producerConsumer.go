package main

import "fmt"

func producer(a chan int) {
	for {
		a <- 1
	}
}

func consumer(a chan int) {
	for {
		b := <-a
		fmt.Println(b)
	}
}

func main() {
	a := make(chan int)
	go producer(a)
	consumer(a)
}
