package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 2
	for j < 20 {
		fmt.Println(j)
		j += 2
	}

	k := 1
	switch k {
	case 1:
		fmt.Println("One One")
		fmt.Println("One Two")
	case 3:
		fmt.Println("One Three")
	case 4:
		fmt.Println("One")
	default:
		fmt.Println("Unknown")
	}
}
