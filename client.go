package marketdata

import (
	"github.com/chentiangang/marketdata/model"

	"github.com/chentiangang/marketdata/dongfang"
)

type Client struct {
	Market         MarketImpl
	Kline          KlineImpl
	RealTimeQuotes RealTimeQuoteImpl
	Quote          QuoteImpl
	Indicator      IndicatorImpl
}

type MarketImpl interface {
	Fetch() ([]model.Quote, error)
}

type KlineImpl interface {
	Fetch(symbol string, period string, limit string) (model.Kline, error)
}

type RealTimeQuoteImpl interface {
	Fetch() chan []model.QuotePtr
	Set(symbols []string)
	Close()
	//Add(symbol, exchange string)
	//Delete(symbol string)
}

type QuoteImpl interface {
	Get(symbol string)
}

type IndicatorImpl interface {
}

func NewDefaultClient() *Client {
	return &Client{
		Market:         dongfang.NewDefaultMarketRequest(),
		Kline:          dongfang.NewDefaultKlineRequest(),
		RealTimeQuotes: dongfang.NewRealtimeQuoteRequest(),
	}
}
