package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/quote"
)

func infostock() {
	flag.Parse()

	if len(os.Args[1]) == 0 {
		log.Fatal("No argument provided. Put some stock symbol. Exemple: %v CLDR", os.Args[1])
	}

	sym := flag.Args()
	iter := quote.List(sym)

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

func startwebserver(tickers []string) {
	http.Handle("/", &graphHandler{tickers})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Unable to start critical service: %v", err)
	}
}

type graphHandler struct {
	tickers []string
}

func (gh *graphHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page := charts.NewPage()

	for _, sym := range gh.tickers {
		line := charts.NewLine()
		line.SetGlobalOptions(charts.TitleOpts{Title: sym})

		x, y := getData(sym)

		line.AddXAxis(x)
		line.AddYAxis(sym, y)
		page.Add(line)
	}
	page.Render(w)
}

func getData(sym string) ([]string, []float64) {
	now := time.Now()
	yearAgo := now.AddDate(-1, 0, 0)

	p := &chart.Params{
		Symbol:   sym,
		Start:    datetime.New(&yearAgo),
		End:      datetime.New(&now),
		Interval: datetime.OneDay,
	}

	iter := chart.Get(p)
	count := iter.Count()

	fmt.Printf("We got %v data points\n", count)

	x := make([]string, count)
	y := make([]float64, count)

	i := 0
	for iter.Next() {
		d := iter.Bar()
		date := time.Unix(int64(d.Timestamp), 0).Format("2006-01-02")
		price, _ := d.Close.Round(2).Float64()

		x[i] = date
		y[i] = price
		i++
	}
	return x, y
}

func main() { // go run . ITUB MSFT AMZN
	flag.Parse()

	if len(os.Args[1]) == 0 {
		log.Fatal("No argument provided. Put some stock symbol. Exemple: %v CLDR", os.Args[1])
	}

	tickers := flag.Args()

	infostock()
	startwebserver(tickers)
}
