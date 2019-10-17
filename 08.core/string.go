package main

import (
	"fmt"
	"strings"
)

func main() {
	strings.ToLower("Hola")
	arr := []byte("abcd")
	fmt.Println(arr) // byte characters
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(str)
}
