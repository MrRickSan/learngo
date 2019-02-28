// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multiple
//
//  1. Declare two variables using
//     multiple variable declaration statement
//
//  2. The first variable's name should be active
//  3. The second variable's name should be delta
//
//  4. Print them all
//
// HINT
//  You should declare a bool and an int variable
//
// EXPECTED OUTPUT
//  false 0
// ---------------------------------------------------------

func main() {
	var (
		active bool
		delta  int
	)

	fmt.Println(active, delta)
}
