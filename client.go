package dongfang

type Client struct {
	Market MarketImpl
	Kline  KlineImpl
}

type MarketImpl interface {
	Get(marketType string) []Stock
}

type KlineImpl interface {
	Get(symbol string) []Kline
}

type QuoteImpl interface {
}
