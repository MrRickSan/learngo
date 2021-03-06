// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: No Conversions Allowed
//
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.
//
// EXPECTED OUTPUT
//  10h0m0s later...
// ---------------------------------------------------------

func main() {
	const later = 10

	// It doesn't work because later was defined as int type
	// becoming incompatible with time.Duration type.
	// Making later typeless it will be converted to time.Duration type automatically.
	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)
}
