package main

import "fmt"

func main() {
	books := [4]string{"Gold in mine", "Crash of 2008", "Albert is Genious"}
	fmt.Printf("%q\n", books)

	// Ellipses determine the length of array auomatically
	animals := [...]string{"Monkey", "Shark", "Dolphin"}
	fmt.Printf("%q\n", animals)

	var old_books [5]string
	old_books[0] = "Dracula"
	old_books[1] = "Island"
	old_books[2] = "1984"

	newBooks := [5]string{"Maquiavel", "Bible"}
	old_books = newBooks

	fmt.Printf("books: %#v\n", books)
	fmt.Printf("new books: %#v\n", newBooks)
}
