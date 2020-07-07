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
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

func main() {
	// use your solution from the previous exercise

	type Item struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	type Game struct {
		Item
		Genre string `json:"genre"`
	}

	videogames := []Game{{
		Item: Item{
			ID:    1,
			Name:  "god of war",
			Price: 50,
		},
		Genre: "action adventure",
	},
		{Item: Item{ID: 2, Name: "x-com 2", Price: 30},
			Genre: "strategy",
		},
		{Item: Item{ID: 3, Name: "minecraft", Price: 20},
			Genre: "sandbox",
		},
	}

	games := make(map[int]Game)
	for _, g := range videogames {
		games[g.ID] = g
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

			g, ok := games[id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n", g.ID, g.Name, "("+g.Genre+")", g.Price)

		case save:
			out, err := json.MarshalIndent(videogames, "", "\t")
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(out))
		case quit, "q":
			fmt.Println("bye!")
			return
		}

	}
}
