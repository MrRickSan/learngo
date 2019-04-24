package main

import "fmt"

func main() {
	color := "red"

	// short circuit
	// the first value is true, go skips the second
	fmt.Println("redish color?",
		color == "red" || color == "dark red")

	color = "dark red"

	fmt.Println("redish color?",
		color == "red" || color == "dark red")

	fmt.Println("greenish color?",
		color == "green" || color == "light green")
}
