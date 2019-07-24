package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1) // Intn->[0, n) means 0 included and n excluded
		fmt.Printf(" %d", n)
	}
}
