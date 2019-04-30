package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number\n")
	}

	// switch {
	// case i > 100:
	// 	fmt.Println("big positive number")
	// case i > 0:
	// 	fmt.Println("positive number")
	// default:
	// 	fmt.Println("number")
	// }
}
