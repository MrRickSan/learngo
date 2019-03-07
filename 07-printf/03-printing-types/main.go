package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet variable type: %T\n", planet)
	fmt.Printf("Planet: %s\n", planet)                     // %s = string
	fmt.Printf("Distance: %d millions of kms\n", distance) // %d = decimal
	fmt.Printf("Orbital Period: %f days\n", orbital)       // %f = float
	fmt.Printf("Does %s have life? %t\n", planet, hasLife) // %t = bool

	// Argument Index (in printf indexes start at 1!)
	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n",
		planet, distance)

	fmt.Println()
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Orbital Period: %.3f days\n", orbital)
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
}
