package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		width  uint8 = 255 // max value
		height       = 255
	)

	width++
	// 256? resets to its minimum value: 0

	if int(width) < height {
		fmt.Println("height is greater")
	}

	fmt.Printf("width: %d height: %d\n", width, height)

	// int8(math.MaxInt8 + 1)
	var n int8 = math.MaxInt8
	fmt.Println("max int8:     ", n)   // 127
	fmt.Println("max int8 + 1 :", n+1) // -128

	n = math.MinInt8
	fmt.Println("min int8:     ", n)   // -128
	fmt.Println("min int8 - 1 :", n-1) // 127

	var un uint8
	fmt.Println("max uint8:     ", un)   // 0
	fmt.Println("max uint8 - 1 :", un-1) // 255

	fmt.Println("max float:		", f*1.2)
}
