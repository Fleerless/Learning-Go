package main

import "fmt"

func main() {
	x:= 51
	y:= 10

	if x<y {
		fmt.Printf("%d is less than %d", x, y)
	} else {
		fmt.Printf("%d is less than %d", y, x)
	}

	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("What color is this?")
	}


	// Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	default: 
	fmt.Println("Color is UNKNOWN")
	}
}