// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare string
//
//  1. Declare a string variable
//
//  2. Print that variable
//
// EXPECTED OUTPUT
//  ""
// ---------------------------------------------------------

func main() {
	var s string
	fmt.Printf("s (%T): %q\n", s, s)

	// %T prints the type of the value
	// %q prints an empty string
}
