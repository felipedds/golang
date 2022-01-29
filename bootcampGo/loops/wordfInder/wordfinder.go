package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == w {
				fmt.Printf("#%-2d: %q\n", i, w)
			}
		}
	}
}

// go run wordfinder.go again
