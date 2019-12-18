package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(removeDuplicates("Porto"))
}

func removeDuplicates(a string) string {
	newString := a

	for i := 0; i < len(a); i++ {
		char := string(a[i])

		n := strings.Count(newString, char)
		if n > 1 {
			newString = strings.ReplaceAll(newString, char, "")
		}

	}
	return newString
}
