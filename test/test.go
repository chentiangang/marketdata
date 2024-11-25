package main

import (
	"fmt"
	"time"

	"github.com/chentiangang/marketdata"
)

func main() {
	client := marketdata.NewDefaultClient()

	client.RealTimeQuotes.SetSymbols([]string{"0.002957", "1.600360"})
	defer client.RealTimeQuotes.Close()

	for data := range client.RealTimeQuotes.Fetch() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), data)
	}

}
