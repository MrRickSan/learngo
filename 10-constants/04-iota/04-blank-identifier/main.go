package main

import "fmt"

// Timezones
// EST -5
// MST -7
// PST -8

func main() {
	const (
		EST = -(5 + iota)
		_   // blank identifier
		MST
		PST
	)

	fmt.Println(EST, MST, PST)
}
