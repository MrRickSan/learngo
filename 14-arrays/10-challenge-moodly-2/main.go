package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("[Your name] [positive|negative]")
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"sad 😞", "terrible 😩", "bad 👎"},
		{"happy 😀", "awesome 😎", "good 👍"},
	}

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(moods[0]))

	var n int
	if mood != "negative" {
		n = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[n][num])
}
