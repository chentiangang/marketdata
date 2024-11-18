package dongfang

type Client struct {
	Market Marketer
	Kline  KlineReq
}

type Marketer interface {
	Get(marketType string) []Stock
}

type KlineReq interface {
	Get(symbol string) []Kline
}
