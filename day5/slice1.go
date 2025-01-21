package main

import (
	"fmt"
)

func accessSlice(s []int, key int) (str1 string) {

	defer func() (str string) {
		if r := recover(); r != nil {
			fmt.Println("paniced and recovered by ashok")
			str1 = fmt.Sprintf("internal error : %v \n", r)

		}
		return str
	}()

	var item int

	item = s[key]

	str1 = fmt.Sprintf("item %d , value %d", key, item)
	return str1

}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	var index int
	fmt.Println("enter the index : ")
	fmt.Scanln(&index)

	str := accessSlice(slice, index)
	fmt.Println(str)
}
