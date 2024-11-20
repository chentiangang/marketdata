package marketdata

import (
	"marketdata/dongfang"
	"marketdata/model"
)

type Client struct {
	Market         MarketImpl
	Kline          KlineImpl
	RealTimeQuotes RealTimeQuoteImpl
	Quote          QuoteImpl
}

type MarketImpl interface {
	Fetch() ([]model.Stock, error)
}

type KlineImpl interface {
	Get(symbol string)
}

type RealTimeQuoteImpl interface {
	Get()
}

type QuoteImpl interface {
	Get()
}

func NewDefaultClient() *Client {
	return &Client{
		Market: dongfang.NewDefaultMarketRequest(),
	}

}
