package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("\nmultiplication:")
	createTable("multiplication", max)

	fmt.Println("\ndivision:")
	createTable("division", max)

	fmt.Println("\nsubtraction:")
	createTable("subtraction", max)

	fmt.Println("\naddition:")
	createTable("addition", max)
}

func createTable(operation string, max int) {
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
		operationType(operation, max, i)

		fmt.Println()
	}
}

func operationType(operation string, max, i int) {

	switch operation {
	case "multiplication":
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
	case "division":
		for j := 0; j <= max; j++ {
			if j == 0 {
				fmt.Printf("%5d", j)
				continue
			}
			fmt.Printf("%5d", i/j)
		}
	case "subtraction":
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i-j)
		}
	case "addition":
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i+j)
		}
	}
}
