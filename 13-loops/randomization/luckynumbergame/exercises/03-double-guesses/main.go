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
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one those numbers.

The greater your number is, the harder it gets.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 || guess2 < 0 {
		fmt.Println("Please pick two positive numbers.")
		return
	}

	greater := guess
	if guess < guess2 {
		greater = guess2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(greater + 1)

		if n == guess || n == guess2 {
			fmt.Println("🎉 YOU WIN! ")
			return
		}
	}

	fmt.Println("☠️ YOU LOST... try again?")
}
