package main

import "fmt"

// Person type
type Person struct {
	firstName string
	lastName  string
}

func (p Person) fullName() string {
	return p.firstName + " ," + p.lastName
}

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Trivolta",
	}

	fmt.Println(p.fullName())
}
