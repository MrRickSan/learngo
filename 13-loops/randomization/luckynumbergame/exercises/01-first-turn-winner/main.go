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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
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

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 1; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("🎉 YOU WIN! ")
			if turn == 1 {
				fmt.Println("AND IT WAS IN THE FIRST TRY! SO LUCKY 🍀 ! CONGRATULATIONS!!!")
			}
			return
		}
	}

	fmt.Println("☠️ YOU LOST... try again?")
}
