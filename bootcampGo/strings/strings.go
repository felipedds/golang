package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Raw strings
func rawString() {
	var s string
	s = "how are you?" // String literal
	fmt.Println(s)
	s = `how are you?` // Raw string literal
	fmt.Println(s)
	s = `
	<html>
		<body>Hello</body>
	<html>` // Raw string literal
	fmt.Println(s)
}

func lengthString() {
	name := "алф"
	fmt.Println(name)
	fmt.Println("Length of string in bytes: ", len(name))
	fmt.Println("Length of string: ", utf8.RuneCountInString(name))
}

func banger() {
	msg := "bunger"
	l := len(msg)
	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s) // convert to upper case
	fmt.Println(s)
}

func main() {
	rawString()
	lengthString()
	banger()
}
