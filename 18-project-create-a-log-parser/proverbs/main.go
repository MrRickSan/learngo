package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// To get input from a file:
	// go run main.go < proverbs.txt
	in := bufio.NewScanner(os.Stdin)

	defer os.Stdin.Close()

	var lines int
	for in.Scan() {
		lines++
		fmt.Println("Scanned text: ", in.Text())
	}

	fmt.Printf("There are %d line(s). \n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR: ", err)
	}
}
