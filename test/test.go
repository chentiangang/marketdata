package main

import (
	"fmt"
	"marketdata"
)

func main() {
	client := marketdata.NewDefaultClient()

	client.RealTimeQuotes.SetSymbols([]string{"0.002957"})
	ch := client.RealTimeQuotes.Fetch()
	defer client.RealTimeQuotes.Close()

	for {
		select {
		case quote := <-ch:
			fmt.Println(quote)
		}
	}
}
