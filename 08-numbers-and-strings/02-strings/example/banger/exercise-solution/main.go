package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := strings.ToUpper(os.Args[1])
	mark := strings.Repeat("!", utf8.RuneCountInString(msg))
	fmt.Println(mark + msg + mark)
}
