package main

import (
	"fmt"
)

func main() {

	arr := [...]string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	var index1, index2 int

	// index1 = 2
	// index2 = 4
	fmt.Scanln(&index1)
	fmt.Scanln(&index2)

	if index1 >= 0 && index1 <= 7 && index2 >= 0 && index2 <= 7 {
		slc1 := arr[0 : index1+1]
		slc2 := arr[index1 : index2+1]
		slc3 := arr[index2:len(arr)]
		fmt.Println(slc1)
		fmt.Println(slc2)
		fmt.Println(slc3)

	} else {
		fmt.Println("Incorrect Indexes")
	}

}
