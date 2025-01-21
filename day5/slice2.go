package main

import (
	"errors"
	"fmt"
)

var errorOutOfIndex = errors.New("length of slice should be more than index")

func accessSlice1(s []int, key int) (str string, err error) {
	if key >= (len(s)) {
		// fmt.Println(errorOutOfIndex)
		return "", errorOutOfIndex
	} else {
		// fmt.Printf("item %d, value %d", key, s[key])
		s := fmt.Sprintf("item %d, value %d", key, s[key])
		return s, nil
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	var index int
	fmt.Println("enter the value of index: ")
	fmt.Scanln(&index)
	str1, er := accessSlice1(s, index)

	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println(str1)
	}

}
