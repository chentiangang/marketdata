package main

import (
	"fmt"

	"github.com/chentiangang/marketdata"
)

func main() {
	mkcli := marketdata.NewDefaultClient()

	markets, err := mkcli.Market.Fetch()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range markets {
		fmt.Println(m)
	}

	//mkcli.RealTimeQuotes.Set([]string{"0.002957", "1.600360"})
	//defer mkcli.RealTimeQuotes.Close()

	//line, err := mkcli.Kline.Fetch("0.002957", "15", "120")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, i := range line.Snapshots {
	//	fmt.Println(i)
	//}
	//
	//for data := range mkcli.RealTimeQuotes.Fetch() {
	//	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), data)
	//}

}
