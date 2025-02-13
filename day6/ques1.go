package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var str string
	fmt.Println("enter the conversation : ")
	fmt.Scanln(&str)

	start := 0

	for j, i := range str {
		if i == '^' {
			break
		} else {

			if i == '#' {
				// fmt.Println("bob", j)
				substr := str[start:j]
				start = j + 1
				ch := make(chan string)
				wg.Add(2)
				go func() {
					defer wg.Done()
					// wg.Add(1)
					ch <- substr
				}()
				go func() {
					defer wg.Done()
					// wg.Add(1)
					fmt.Printf("\n bob : %v,", <-ch)
				}()

				wg.Wait()
			}
			if i == '$' {
				// fmt.Println("alice", j)

				substr := str[start:j]
				start = j + 1
				ch := make(chan string)
				wg.Add(2)
				go func() {
					defer wg.Done()

					ch <- substr
				}()

				go func() {
					defer wg.Done()
					// wg.Add(1)
					fmt.Printf("\n alice : %v,", <-ch)
				}()

				wg.Wait()
			}
		}
	}
	fmt.Println()
	// fmt.Println("ashoka")
}

// "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
