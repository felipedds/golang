package main

import (
	"fmt"
	"os"
	"strings"
)

// Switch statement
func citycontry() {
	city := strings.ToLower(os.Args[1])

	switch city {
	case "paris":
		fmt.Println("France")
	case "tokyo":
		fmt.Println("Japan")
	case "sucre", "la paz":
		fmt.Println("Bolivia")
	default:
		fmt.Println("Where?")
	}
}

func positivenegative() {

	switch i := 100; {
	case i > 50:
		fmt.Println("Big")
		fallthrough
	case i > 0:
		fmt.Println("Positive")
	case i < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("0")
	}
}

func main() {
	citycontry()
	positivenegative()
}
