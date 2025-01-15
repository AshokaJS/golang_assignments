package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "monday"
	m[2] = "tuesday"
	m[3] = "wednesday"
	m[4] = "thursday"
	m[5] = "friday"
	m[6] = "saturday"
	m[7] = "sunday"

	var index int

	fmt.Scanln(&index)

	if index >= 1 && index <= 7 {
		fmt.Println(m[index])
	} else {
		fmt.Println("Not a day")
	}

}
