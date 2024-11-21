package model

type Quote struct {
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	DifferenceValue  float64 `json:"price"`
	PriceLimit       float64 `json:"price_limit"`
	TurnoverRate     float64 `json:"turnover_rate"`     // 换手率
	TotalValue       int64   `json:"total_value"`       // 总市值
	CirculatingValue int64   `json:"circulating_value"` // 流通
	TotalShares      int64   `json:"total_shares"`

	Exchange int    `msgpack:",omitempty"` // 交易所
	Industry string `msgpack:",omitempty"` // 行业
	Alias    string `msgpack:",omitempty"` //
}

type QuotePtr struct {
	Symbol           string   `json:"symbol"`
	Name             string   `json:"name"`
	Price            *float64 `json:"price"`             // 价格
	DifferenceValue  *float64 `json:"price"`             // 涨跌额
	PriceLimit       *float64 `json:"price_limit"`       // 涨跌幅
	TurnoverRate     *float64 `json:"turnover_rate"`     // 换手率
	TotalValue       *int64   `json:"total_value"`       // 总市值
	CirculatingValue *int64   `json:"circulating_value"` // 流通市值
	TotalShares      *int64   `json:"total_shares"`

	Exchange int    `msgpack:",omitempty"` // 交易所
	Industry string `msgpack:",omitempty"` // 行业
	Alias    string `msgpack:",omitempty"` //
}
