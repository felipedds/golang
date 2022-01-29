package main

import "fmt"

func main() {

	students := [2][3]int{
		{5, 6, 1},
		{2, 8, 7},
	}

	employees := [...][3]int{
		{10, 6, 4},
		{9, 8, 3},
	}

	fmt.Printf("%v", students)

	for _, v := range employees {
		fmt.Printf("%d", v)
	}
}
