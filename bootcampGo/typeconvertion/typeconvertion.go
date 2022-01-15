package main

import "fmt"

func main() {
	speed := 100 // int
	force := 2.5 // float64

	speed = int(float64(speed) * force)
	fmt.Println(speed)
}
