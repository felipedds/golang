package main

import (
	"fmt"
	"path"
)

func path_separator() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}

func discard() {
	var file string
	_, file = path.Split("css/main.css")
	fmt.Println("file: ", file)
}

func short_declaration() {
	dir, file := path.Split("css/main.css")
	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}

func main() {
	path_separator()
	discard()
	short_declaration()
}
