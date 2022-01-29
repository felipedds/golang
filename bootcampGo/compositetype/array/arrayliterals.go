package main

import "fmt"

func main() {
	books := [4]string{"Gold in mine", "Crash of 2008", "Albert is Genious"}
	fmt.Printf("%q\n", books)

	// Ellipses determine the length of array auomatically
	animals := [...]string{"Monkey", "Shark", "Dolphin"}
	fmt.Printf("%q\n", animals)
}
