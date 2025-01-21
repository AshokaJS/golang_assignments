package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// s := "string"
	var s string
	fmt.Println("enter the string which you want to reverse : ")
	fmt.Scanln(&s)
	ch := make(chan string) // channel for transferring values between first go routine to second like reversed string from first go routine to second go routine.
	wg.Add(1)
	go func() {
		var v int = len(s)
		wg.Add(1)
		defer wg.Done()

		sbyte := []rune(s)

		for i := 0; i < v/2; i++ {

			sbyte[i], sbyte[v-1-i] = sbyte[v-1-i], sbyte[i]

		}

		s = string(sbyte)

		ch <- s

	}()

	go func() {
		defer wg.Done()
		fmt.Println("reversed string is : ")
		fmt.Println(<-ch)
	}()

	wg.Wait()
	// time.Sleep(2 * time.Second)
	// fmt.Println("string reversed")

}
