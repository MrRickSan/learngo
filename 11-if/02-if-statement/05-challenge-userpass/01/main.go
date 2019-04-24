package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOk = "Access granted to %q.\n"

	user = "rick"
	pass = "1234"
)

func main() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l != 2 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user {
		fmt.Printf(errUser, u)
	} else if p != pass {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOk, u)
	}
}
