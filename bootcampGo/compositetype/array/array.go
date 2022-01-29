package main

import "fmt"

func main() {
	var books [4]string

	books[0] = "Kafka's revenge"
	books[1] = "Stay golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " Second Edition"

	fmt.Printf("books :%T\n", books)
	fmt.Println("books :", books)
	fmt.Printf("books : %#q\n", books)
	fmt.Printf("books : %#v\n", books)

	var (
		sbooks [4]string
	)

	for i := range sbooks {
		sbooks[i] = books[i]
		fmt.Printf("books : %q", sbooks[i])
	}

	for _, v := range sbooks {
		v += " won't effect."
		fmt.Printf(v)
	}
}
