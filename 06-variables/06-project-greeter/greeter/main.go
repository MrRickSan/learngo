package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello great", os.Args[1], "!")
	fmt.Println("Hey!", os.Args[2], "!")
	fmt.Println("Welcome", os.Args[3], "!")

	name, age := "bilbo", 2019
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("And I love adventures!")

}
