package main

import "fmt"

func main() {
	// make sliced array
	slice1 := make([]float64, 5)

	slice2 := append(slice1, 20, 12, 23, 34, 45)
	copyslice := make([]float64, 3)
	copy(copyslice, slice2)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(copyslice)
}
