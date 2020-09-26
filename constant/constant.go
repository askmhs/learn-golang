package main

import "fmt"

func main() {
	/**
	 * Constant will error if the value was changed
	 * and constant must be initialized
	 */
	const firstName string = "Muhammad Harits"
	const lastName string = "Syaifulloh"

	fmt.Println(firstName + " " + lastName)
}
