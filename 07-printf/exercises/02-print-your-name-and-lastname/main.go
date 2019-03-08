// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	firstName, lastName := "Ricardo", "Oliveira"
	fmt.Printf("My name is %s and my lastname is %s\n", firstName, lastName)

	// BONUS
	msg := "My name is %s and my lastname is %s\n"
	fmt.Printf(msg, firstName, lastName)
}
