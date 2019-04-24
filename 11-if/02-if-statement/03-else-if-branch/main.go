package main

import "fmt"

func main() {
	score, valid := 3, true

	if score > 3 && valid {
		fmt.Println("good")
	} else if score == 2 && !valid {
		fmt.Println("meh...")
	} else {
		fmt.Println("low")
	}
}
