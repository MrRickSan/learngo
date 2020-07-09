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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func main() {
	// use your solution from the previous exercise

	type item struct {
		id    int    `json:"id"`
		name  string `json:"name"`
		price int    `json:"price"`
	}

	type game struct {
		item
		genre string `json:"genre"`
	}

	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
	}

	var videogames []jsonGame
	err := json.Unmarshal([]byte(data), &videogames)
	if err != nil {
		fmt.Printf("error during unmarshal operation: %s\n", err)
		return
	}

	var games []game
	for _, vg := range videogames {
		games = append(games, game{item{vg.ID, vg.Name, vg.Price}, vg.Genre})
	}

	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	fmt.Printf("Rick's game store has %d games.\n", len(videogames))

	const (
		list = "list"
		id   = "id"
		quit = "quit"
		save = "save"
	)

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Printf("  > %s : lists all the games\n", list)
		fmt.Printf("  > %s : queries a game by id\n", id+" N")
		fmt.Printf("  > %s : exports the data to json and quits\n", save)
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
				fmt.Printf("#%d: %-15q %-20s $%d\n", g.ID, g.Name, "("+g.Genre+")", g.Price)
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

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)

		case save:
			out, err := json.MarshalIndent(videogames, "", "\t")
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(out))

			return
		case quit, "q":
			fmt.Println("bye!")
			return
		}

	}
}
