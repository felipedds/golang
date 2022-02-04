package main

import "fmt"

// Describe the indexes manually
func main() {

	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,  // ethereum
		WAN: 120.5, // wanchain
	}

	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])
}
