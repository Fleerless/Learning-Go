package main

import "fmt"

func main() {
	// Define Map
	emails:= make(map[string]string)
	
	// assign Key Values

	emails["Bob"] = "bob@gmail.com"
	emails["stacy"] = "stacy@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails["Bob"])
	fmt.Println(emails)

	delete(emails, "Bob")
	fmt.Println(emails)

}