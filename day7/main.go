package main

import (
	"fmt"
	"sync"
	"time"
)

var mg sync.Mutex

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	go func() {
		mg.Lock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
		mg.Unlock()
	}()

	go func() {
		mg.Lock()
		n++
		mg.Unlock()
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
