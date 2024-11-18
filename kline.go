package dongfang

type Kline struct {
	Timestamp   time.Time `json:"timestamp"`    // 时间戳
	Open        float64   `json:"open"`         // 开盘价
	Close       float64   `json:"close"`        // 收盘价
	High        float64   `json:"high"`         // 最高价
	Low         float64   `json:"low"`          // 最低价
	Volume      int64     `json:"volume"`       // 成交量
	Amount      float64   `json:"amount"`       // 成交额
	ChangePct   float64   `json:"change_pct"`   // 涨跌幅
	ChangeAmt   float64   `json:"change_amt"`   // 涨跌额
	Amplitude   float64   `json:"amplitude"`    // 振幅
	TurnoverPct float64   `json:"turnover_pct"` // 换手率
}

type Result struct {
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	Klines []Kline `json:"kline"`
}
