package main

import "fmt"

func main() {
	// // Arrays
	// var fruitArr [2]string

	// // Assign Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// SLICES
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}