package main

import f "fmt"

var enabled bool

func main() {
	// var enabled bool
	hey()
	enabled = true
	f.Println(enabled)
	f.Println("Hello!")
	bye()
}
