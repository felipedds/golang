package main

import (
	"fmt"
)

func append_builtin() {
	units := []int{1, 2, 3, 4, 5}
	fmt.Printf("units: %+v\n", units)
	dozens := []int{11, 12, 13, 14, 15}
	units = append(units, dozens...) // Append
	fmt.Printf("units + dozens: %+v\n", units)

	nums := []int{9, 7, 5}
	evens := append(nums, []int{2, 4, 6}...)
	fmt.Println(nums, evens)
}

func slices() {
	games := []string{"Pokemon", "Star Wars", " Flinstons", "Simpsons", "Rick and Morty"}
	newGames := games[1:3] // Slice

	fmt.Printf("games: %#v\n", games)
	fmt.Printf("games len: %#v\n", len(games))
	fmt.Printf("new games: %#v\n", newGames)
	fmt.Printf("new games len: %#v\n", len(newGames))
}

func main() {
	slices()
	append_builtin()
}
