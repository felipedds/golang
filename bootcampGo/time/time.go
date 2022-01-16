package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func partsofday() {
	switch h := time.Now().Hour(); {
	case h > 18:
		println("Good evening")
	case h >= 12:
		println("Good aternoon")
	case h >= 6:
		println("Good morning")
	default:
		fmt.Println("Good morning")
	}
	fmt.Printf("%d:%d:%d\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}

func seasons() {
	if len(os.Args) != 2 {
		fmt.Println("Gimme a month name.")
	}

	switch m := strings.ToLower(os.Args[1]); m {
	case "december", "january", "february":
		fmt.Println("Winter")
	case "march", "april", "may":
		fmt.Println("Spring")
	case "june", "july", "august":
		fmt.Println("Summer")
	case "september", "october", "november":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month \n", m)
	}
}

func main() {
	seasons()
	partsofday()
}
