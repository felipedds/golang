package main

import (
	"fmt"
	"os"
	"strconv"
)

// nil value means that the value is not initialized yet
func errorhandling() {
	age := os.Args[1]

	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Success converted %q to %d", age, n)
}

func meterstofeet() {
	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("Error: %q", err)
	}

	meter := feet * 0.3048
	fmt.Printf("%g feet is %g meters.\n", feet, meter)
}

func main() {
	meterstofeet()
}
