package main

import "fmt"

func main() {

	for i := 1; i < 100; i++ {
		var output string

		if i%3 == 0 {
			output += "Fizz"
		}

		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			fmt.Println(i)
			continue
		}

		fmt.Println(output)
	}
}
