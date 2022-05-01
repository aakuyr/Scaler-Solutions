package main

import (
	"fmt"
	"json"
)

type Message struct {
	Name, Message string
}

func main() {
	data := []byte(`{"status" : "ok"}`)
	var m Message
	err := json.Unmarshall(data, &m)
	fmt.Println(err)
}
