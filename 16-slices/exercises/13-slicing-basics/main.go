package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

var nums, evens, odds []int

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"

	stringSlice := strings.Fields(data)

	// convert the string slice to int slice
	for _, v := range stringSlice {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}

		nums = append(nums, n)
	}

	fmt.Println("nums:", nums)

	// get evens and odds from nums slice
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		} else {
			odds = append(odds, n)
		}

	}

	fmt.Println("evens:", evens)
	fmt.Println("odds:", odds)
	fmt.Println("middle:", nums[len(nums)/2-1:len(nums)/2+1])
	fmt.Println("first 2:", nums[:2])
	fmt.Println("last 2:", nums[len(nums)-2:])
	fmt.Println("evens last 1:", evens[len(evens)-1:])
	fmt.Println("odds last 2:", odds[len(odds)-2:])
}
