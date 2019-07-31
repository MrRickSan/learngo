package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	names := [...]string{"Luffy", "Saitama", "Tanjiro"}
	distances := [...]int{3, 5, 13, 22, 43}
	data := [...]byte{79, 71, 69, 73, 76}
	ratios := [...]float64{.23}
	alives := [...]bool{false, true, true, false}

	// when don't assign elements is better this way:
	var zero []byte

	fmt.Println("For loops:")
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}
	fmt.Println()
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}
	fmt.Println()
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}
	fmt.Println()
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}
	fmt.Println()
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Println()
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %v\n", i, zero[i])
	}
	fmt.Println()

	fmt.Println("For range loops:")
	for i := range names {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}
	fmt.Println()
	for i := range distances {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}
	fmt.Println()
	for i := range data {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}
	fmt.Println()
	for i := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}
	fmt.Println()
	for i := range alives {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Println()

	for i := range zero {
		fmt.Printf("zero[%d]: %v\n", i, zero[i])
	}
}
