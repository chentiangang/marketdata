package main

import (
	"fmt"
	"marketdata"

	"github.com/chentiangang/xlog"
)

func main() {
	client := marketdata.NewDefaultClient()
	s, err := client.Market.Fetch()
	if err != nil {
		xlog.Error("%s", err)
		return
	}
	fmt.Println(len(s))
}
