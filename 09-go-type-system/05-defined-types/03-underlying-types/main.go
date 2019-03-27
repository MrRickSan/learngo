package main

import (
	"fmt"

	"github.com/mrricksan/learngo/09-go-type-system/05-defined-types/03-underlying-types/weights"
)

func main() {
	type (
		// Gram underlying type is int
		Gram int

		// Kilogram underlying type is int
		Kilogram Gram

		// Ton underlying type is int
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Gram(int(apples)))

	salt = Gram(weights.Gram(100))
	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
	fmt.Printf("Type of main.Gram: %T\n", Gram(1))

	type myGram weights.Gram

	var pepper myGram = 20
	fmt.Printf("Type of pepper: %T\n", pepper)

	pepper = myGram(salt)
}
