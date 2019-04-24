package main

import "fmt"

func main() {
	speed := 100

	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	speed = 20

	// short circuit
	// the first value is false, go skips the second
	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	speed = 50

	fmt.Println("within limits?", speed >= 40 && speed <= 55)

	fmt.Println(1 == 1 && 2 == 2)
}
