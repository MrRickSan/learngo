// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a search word.")
		return
	}
	query := strings.ToLower(args[0])

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		line := in.Text()
		if strings.Contains(strings.ToLower(line), query) {
			fmt.Println(line)
		}
	}
}
