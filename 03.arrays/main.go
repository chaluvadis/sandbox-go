package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)

	var y [5]float64
	fmt.Println(y)

	var z [5]string
	fmt.Println(z)

	var b [5]bool
	fmt.Println(b)

	// Assign values to each array
	x1 := [5]int{
		12, 13, 14, 145, 16,
	}

	for i, value := range x1 {
		fmt.Printf("i = %v, value = %v \n", i, value)
	}
}
