package marketdata

type Client struct {
	Market         MarketImpl
	Kline          KlineImpl
	RealTimeQuotes RealTimeQuoteImpl
	Quote          QuoteImpl
}

type MarketImpl interface {
	Get(marketType string) []Stock
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
	return &Client{}
}
