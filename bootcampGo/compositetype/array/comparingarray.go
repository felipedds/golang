package main

import "fmt"

// Array are comparable when their types are identical
func main() {
	var (
		blue   = [3]int{6, 9, 3}
		red    = [3]int{6, 9, 3}
		yellow = [3]int{2, 7, 5}
		green  = yellow
	)

	fmt.Printf("blue bookcase: %v\n", blue)
	fmt.Printf("red bookcase: %v\n", red)
	fmt.Println("Are they equal?", blue == red)

	fmt.Printf("blue bookcase: %v\n", blue)
	fmt.Printf("red bookcase: %v\n", yellow)
	fmt.Println("Are they equal?", blue == yellow)
	fmt.Printf("green bookcase: %v\n", green)

}
