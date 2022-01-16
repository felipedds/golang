package main

import (
	"fmt"
	"os"
	"strings"
)

func basicfor() {
	var sum int

	for i := 0; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func breakfor() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}
		sum += i
		fmt.Println(i, "->", sum)
		i++
	}
	fmt.Println(sum)
}

func nestedfor() {
	fmt.Printf("%5s", "X")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= 5; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= 5; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

func stringfor() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%q\n", os.Args[i])
	}

	words := strings.Fields("lazy cat jumps and again and again")

	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
		fmt.Printf("%q\n", words[j])
	}
}

func rangefor() {
	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}

	words := strings.Fields("learn the easy way")
	for i, v := range words {
		fmt.Printf("%-2d: %q\n", i+1, v)
		fmt.Printf("%q\n", v)
	}
}

func main() {
	// go run for.go hello cats and monkeys
	basicfor()
	breakfor()
	nestedfor()
	stringfor()
	rangefor()
}
