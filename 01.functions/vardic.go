package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
	mul := func(args ...int) int {
		total := 1
		for _, v := range args {
			total *= v
		}
		return total
	}
	fmt.Println(mul(1, 2, 3, 4, 5))
}
