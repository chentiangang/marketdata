package main

import (
	"fmt"
	"marketdata"

	"github.com/chentiangang/xlog"
)

func main() {
	client := marketdata.NewDefaultClient()
	line, err := client.Kline.Fetch("0.002957", "1", "360")
	if err != nil {
		xlog.Error("%s", err)
		return
	}
	//fmt.Println(line)
	for _, i := range line.Snapshots {
		fmt.Println(i)
	}
}
