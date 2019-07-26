// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	var (
		mode    string
		guess   int
		verbose bool
		err     error
	)

	switch l := len(args); {
	case l < 1:
		fmt.Printf(usage, maxTurns)
		return
	case l == 1:
		guess, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
	case l > 1:
		mode = args[0]
		guess, err = strconv.Atoi(args[1])
		if err != nil || mode != "-v" {
			fmt.Println("Usage: [-v] [NUMBER]")
			return
		}
		verbose = true
	}

	if guess < 0 {
		fmt.Println("Please pick positive numbers.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
