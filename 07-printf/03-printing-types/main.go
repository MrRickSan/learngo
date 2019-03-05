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
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions of kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// Argument Index (in printf indexes start at 1!)
	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n",
		planet, distance)
}
