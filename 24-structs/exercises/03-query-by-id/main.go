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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

func main() {
	// use your solution from the previous exercise

	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	videogames := []game{{
		item: item{
			id:    1,
			name:  "god of war",
			price: 50,
		},
		genre: "action adventure",
	},
		{item: item{id: 2, name: "x-com 2", price: 30},
			genre: "strategy",
		},
		{item: item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	games := make(map[int]game)
	for _, g := range videogames {
		games[g.id] = g
	}

	fmt.Printf("Rick's game store has %d games.\n", len(videogames))

	const (
		list = "list"
		id   = "id"
		quit = "quit"
	)

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Printf("  > %s : lists all the games\n", list)
		fmt.Printf("  > %s : queries a game by id\n", id+" N")
		fmt.Printf("  > %s : quits\n\n", quit)

		if !in.Scan() {
			break
		}

		fmt.Println()

		option := strings.Fields(in.Text())
		if len(option) == 0 {
			continue
		}

		switch option[0] {
		case list:
			for _, g := range videogames {
				fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
			}
		case id:
			if len(option) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(option[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := games[id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)

		case quit, "q":
			fmt.Println("bye!")
			return
		}

	}
}
