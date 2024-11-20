package marketdata

import "project/stock/marketdata/dongfang"

type Client struct {
	Market         MarketImpl
	Kline          KlineImpl
	RealTimeQuotes RealTimeQuoteImpl
	Quote          QuoteImpl
}

type MarketImpl interface {
	Fetch() ([]Stock, error)
}

type KlineImpl interface {
	Get(symbol string) []Kline
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
