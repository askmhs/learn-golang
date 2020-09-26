package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Harits S"
	fmt.Println(name)

	name = "Harits"
	fmt.Println(name)

	newName := "Muhammad Harits Syaifulloh"
	fmt.Println(newName)

	var (
		firstName string = "Muhammad Harits"
		lastName  string = "Syaifulloh"
	)

	fmt.Println(firstName + " " + lastName)
}
