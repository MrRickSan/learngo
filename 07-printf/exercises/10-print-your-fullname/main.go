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
)

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func main() {
	fmt.Printf("Your name is %s and your lastname is %s.\n", os.Args[1], os.Args[2])

	// BONUS: Use a variable for the format specifier
	msg := "Your name is %s and your lastname is %s.\n"
	fmt.Printf(msg, os.Args[1], os.Args[2])
}
