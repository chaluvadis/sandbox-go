package main

import (
	"fmt"
	"strings"
)

// Pass by value
func welcome(name string) string {
	return `Welcome, ` + name
}

// Pointer functions
func updateMessage(msg *string) {
	*msg = strings.Title(*msg)
}

// return multiple values
func returnBoth(msg string) (string, string) {
	wel := welcome(msg)
	upd := welcome(msg)
	updateMessage(&upd)
	return wel, upd
}

func main() {
	msg, upd := returnBoth("surendra")
	fmt.Println(msg, upd)
}
