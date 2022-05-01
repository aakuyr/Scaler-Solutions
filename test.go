package main

import (
	"fmt"
	"net/http"
	"sync"
)

func request(i int, w *sync.WaitGroup) {
	defer w.Done()
	http.Get("https://google.com")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go request(i, &wg)
	}
	wg.Wait()
	fmt.Println("Done")
}
