package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/piquette/finance-go/quote"
)

func stocks() {
	flag.Parse()

	if len(os.Args[1]) == 0 {
		log.Fatal("No argument provided. Put some stock symbol. Exemple: %v CLDR", os.Args[1])
	}

	symbols := flag.Args()
	iter := quote.List(symbols)

	for iter.Next() {
		q := iter.Quote()
		fmt.Printf("%v - %v \n", q.ShortName, q.Symbol)
		fmt.Printf("Current Price $%v \n", q.Ask)
		fmt.Printf("Fifty Day Average $%v \n", q.FiftyDayAverage)
		fmt.Printf("Fifty Two Week Low $%v \n", q.FiftyTwoWeekLow)
		fmt.Printf("Fifty Two Week High $%v \n", q.FiftyTwoWeekHigh)
		fmt.Printf("\n")
	}
}

func main() {
	stocks() // go run . itub msft amzn
}
