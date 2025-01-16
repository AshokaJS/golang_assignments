package main

import (
	"fmt"
)

func main() {
	var roman string
	_, err := fmt.Scanln(&roman)
	if err != nil {
		fmt.Printf("error while reading input, err: %s", err)
	}

	m := make(map[byte]int)

	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	l := len(roman)

	var digit int
	digit = 0

	var i,j int

	if l==1 {
		digit=m[roman[i]]
	}else{
	
	
		for i = 0; i < l-1; i++ {
			j = i + 1

			if m[roman[i]] < m[roman[j]] {
				digit += (m[roman[j]] - m[roman[i]])
				i++
			} else {
				digit += m[roman[i]]
				if j+1 >= l {
					digit += m[roman[j]]
				}
			}

		}
	}


	fmt.Println(digit)

}
