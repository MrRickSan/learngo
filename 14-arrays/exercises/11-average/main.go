// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]
	if len(args) == 0 || len(args) > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var (
		numbers [5]float64
		count   float64
		sum     float64
	)

	for i, v := range args {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		numbers[i] = num
		sum += num
		count++
	}

	fmt.Println("Your numbers:", numbers)
	fmt.Println("Average:", sum/count)
}
