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

	user, user2 = "rick", "jack"
	pass, pass2 = "1234", "1888"
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

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p != pass {
		fmt.Printf(errPwd, u)
	} else if u == user2 && p != pass2 {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOk, u)
	}
}
