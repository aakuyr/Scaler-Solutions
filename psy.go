package main

import "fmt"

func say(message string) {
	fmt.Println(message)
}

func main() {
	go say("Hello")
	go say("World")
}
