package main

import "fmt"

//IOTA ia built-on constant generator which generates ever increasing number
func iotaConst() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

func timezones() {

	// EST = -5 / MST = -7 / PST = -8
	const (
		EST = -(5 + iota)
		_   // blank identifier
		MST
		PST
	)
	fmt.Println(EST, MST, PST)
}

func months() {
	const (
		January = iota + 1
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
	fmt.Println(January, February, March, April, May, June, July, August, September, October, November, December)
}
func seasons() {
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)
	fmt.Println(Spring, Summer, Fall, Winter)
}

func main() {
	iotaConst()
	timezones()
	months()
	seasons()
}
