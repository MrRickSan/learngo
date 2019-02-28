// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Wrong doer
//
//  1. Print a variable
//
//  2. Then declare it
//  (This means: Try to print it before its declaration)
//
//  3. Observe the error
// ---------------------------------------------------------

func main() {
	fmt.Println(age)

	var age int
}
