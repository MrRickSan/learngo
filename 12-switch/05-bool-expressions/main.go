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
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
