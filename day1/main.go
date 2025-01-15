// code for calculating simple interest in golang.
package main

import "fmt"

func main() {
	var principle uint
	var rate uint
	var time uint

	fmt.Scanln(&principle)
	fmt.Scanln(&rate)
	fmt.Scanln(&time)

	var simple_interest uint
	simple_interest = (principle * rate * time) / 100

	fmt.Printf("simple interest of principle amount %v is : %v \n", principle, simple_interest)
}
