package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string
	fmt.Scanln(&st)

	// st = "Europe is far but the US is farther"

	slc := make([]string, 0, 0)

	slice := strings.Split(st, " ")

	m := make(map[string]int)

	var i int

	for i = 0; i < len(slice); i++ {
		m[slice[i]]++
	}

	var max_frequency int
	max_frequency = 0

	for _, i := range m {
		if i > max_frequency {
			max_frequency = i
		}
	}

	mp := make(map[string]int)

	for i = 0; i < len(slice); i++ {
		if m[slice[i]] == max_frequency && mp[slice[i]] == 0 {
			mp[slice[i]]++
			slc = append(slc, slice[i])

		}
	}

	fmt.Println(slc)
}
