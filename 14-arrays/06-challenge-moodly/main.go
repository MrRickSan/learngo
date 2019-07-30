package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("[Your name]")
		return
	}

	name := args[0]

	moods := [...]string{
		"sad 😞", "terrible 😩", "bad 👎",
		"happy 😀", "awesome 😎", "good 👍",
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[num])
}
