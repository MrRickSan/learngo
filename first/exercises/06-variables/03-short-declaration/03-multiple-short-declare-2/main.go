// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multiple Short Declare #2
//
// 	1. Declare two variables using short declaration
//
//  2. `a` variable's value should be 42
//  3. `c` variable's value should be "good"
//
// EXPECTED OUTPUT
//  42 good
// ---------------------------------------------------------

func main() {
	a, c := 42, "good"

	fmt.Println(a, c)
}
