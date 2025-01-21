package main

import (
	"fmt"
	"log"
)

type hello uint64

// type data interface{}

func AcceptAnything(d interface{}) {
	switch v := d.(type) {
	case int:
		fmt.Println("Integer :", v)

	case string:
		fmt.Println("string :", v)

	case bool:
		fmt.Println("boolean :", v)
	case hello:
		fmt.Println("hello :", v)

	}
}

func main() {
	var input int

	if _, err := fmt.Scanln(&input); err != nil {
		log.Fatal(err)
	}

	switch input {
	case 1:
		AcceptAnything(1)
	case 2:
		AcceptAnything("hello world")
	case 3:
		AcceptAnything(true)
	case 4:
		var i hello
		AcceptAnything(i)
	}

}
