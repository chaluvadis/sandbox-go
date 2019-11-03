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

// func main() {
// 	msg, upd := returnBoth("surendra")
// 	fmt.Println(msg, upd)
// }

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}
