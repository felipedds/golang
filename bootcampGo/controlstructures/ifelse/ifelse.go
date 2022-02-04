package main

import (
	"fmt"
	"os"
	"strconv"
)

// If Statement
func ifstatement() {
	score, valid := 3, true

	if score >= 5 && valid {
		fmt.Println("Good")
	} else if score >= 3 && !valid {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}

func validatePassword() {
	args := os.Args
	u, p := args[1], args[2] // User & Password

	if u == "admin" && p == "admin" {
		fmt.Printf("Correct %q", u)
	} else if u == "admin" && p != "admin" {
		fmt.Printf("Wrong Password %q", u)
	} else if u != "admin" && p == "admin" {
		fmt.Printf("Wrong Username %q", u)
	}
}

// Simple Statement allow you to execute a statement inside another statement
func simpleStatement() {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("There was no error, n is", n)
	}

	if a := os.Args; len(a) != 2 {
		fmt.Printf("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q \n", a[1])
	} else {
		fmt.Printf("%s * 2 = %d \n", a[1], n*2)
	}
}

func main() {
	simpleStatement()
}
