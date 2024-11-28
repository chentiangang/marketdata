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


func (s *Quote) Join() string {
	return strings.Join([]string{s.Symbol, s.Name, s.Industry, s.Alias}, "-")
}

func (s *Quote) Update(sptr QuotePtr) {
	if sptr.Price != nil {
		s.Price = *sptr.Price
		if s.TotalShares != 0 {
			s.TotalValue = int64(s.Price) * s.TotalShares
		}
	}

	if sptr.PriceLimit != nil {
		s.PriceLimit = *sptr.PriceLimit
	}

	if sptr.TurnoverRate != nil {
		s.TurnoverRate = *sptr.TurnoverRate
	}

	if sptr.DifferenceValue != nil {
		s.DifferenceValue = *sptr.DifferenceValue
	}

	if sptr.TotalValue != nil {
		s.TotalValue = *sptr.TotalValue
	}

	if sptr.CirculatingValue != nil {
		s.CirculatingValue = *sptr.CirculatingValue
	}
	if sptr.TotalShares != nil {
		s.TotalShares = *sptr.TotalShares
	}
}

