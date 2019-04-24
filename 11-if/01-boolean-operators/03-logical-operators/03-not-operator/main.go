package main

import "fmt"

func main() {
	var on bool

	on = !on
	fmt.Println(on)

	on = !!on // in a real code, you should never need to do this
	fmt.Println(on)
}
