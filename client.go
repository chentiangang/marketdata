package marketdata

import (
	"marketdata/dongfang"
	"marketdata/model"
	"net/http"
)

type Client struct {
	Market         MarketImpl
	Kline          KlineImpl
	RealTimeQuotes RealTimeQuoteImpl
	Quote          QuoteImpl
}

type MarketImpl interface {
	Fetch() ([]model.Quote, error)
}

type KlineImpl interface {
	Fetch(symbol string, period string, limit string) (model.Kline, error)
}

type RealTimeQuoteImpl interface {
	Request(symbol string)
	Parse(*http.Response) ([]model.QuotePtr, error)
}

type QuoteImpl interface {
	Get(symbol string)
}

func NewDefaultClient() *Client {
	return &Client{
		Market: dongfang.NewDefaultMarketRequest(),
		Kline:  dongfang.NewDefaultKlineRequest(),
	}
}
