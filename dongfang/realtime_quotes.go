package dongfang

import "net/http"

type RealtimeQuoteRequest struct {
	BaseURL string
	Request *http.Request
}

type RealtimeQuoteResponse struct {
	Rc     int                `json:"rc"`
	Rt     int                `json:"rt"`
	Svr    int64              `json:"svr"`
	Lt     int                `json:"lt"`
	Full   int                `json:"full"`
	Dlmkts string             `json:"dlmkts,omitempty"`
	Data   *RealtimeQuoteData `json:"data,omitempty"`
}

type RealtimeQuoteData struct {
	Total int `json:"total"`
	Diff  map[string]struct {
		//F1   int     `json:"f1,omitempty"`
		F2 *int `json:"f2,omitempty"`
		F3 *int `json:"f3,omitempty"` // 涨跌幅
		F4 *int `json:"f4,omitempty"`
		//F5   int     `json:"f5,omitempty"`
		//F6   float64 `json:"f6,omitempty"`
		//F7   int     `json:"f7,omitempty"`
		F8 *int `json:"f8,omitempty"` //换手率
		//F9   int     `json:"f9,omitempty"`
		//F10  int     `json:"f10,omitempty"`
		F12 string `json:"f12,omitempty"`
		//F13  int     `json:"f13,omitempty"`
		F14 string `json:"f14,omitempty"`
		//F18  int     `json:"f18,omitempty"`
		//F19  int     `json:"f19,omitempty"`
		F20 *int `json:"f20,omitempty"`
		F21 *int `json:"f21,omitempty"`
		F84 *int `json:"f84,omitempty"`
		//F22  int     `json:"f22,omitempty"`
		//F30  int     `json:"f30,omitempty"`
		//F31  int     `json:"f31,omitempty"`
		//F32  int     `json:"f32,omitempty"`
		//F100 string  `json:"f100,omitempty"`
		//F112 float64 `json:"f112,omitempty"`
		//F125 int     `json:"f125,omitempty"`
		//F139 int     `json:"f139,omitempty"`
		//F148 int     `json:"f148,omitempty"`
		//F152 int     `json:"f152,omitempty"`
	} `json:"diff,omitempty"`
}

var defaultRealTimeQuoteHeaders = map[string]string{
	"Content-type":    "text/event-stream",
	"Accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"Cache-Control":   "no-cache",
	"Connection":      "keep-alive",
	//"Host", sse.Host)
	"Origin":             "https://quote.eastmoney.com",
	"Referer":            "https://quote.eastmoney.com/zixuan/?from=home",
	"Sec-ch-ua":          `"Chromium";v="128", "Not;A=Brand";v="24", "Microsoft Edge";v="128"`,
	"Sec-ch-ua-mobile":   "?0",
	"Sec-ch-ua-platform": "macOS",
	"Sec-fetch-dest":     "empty",
	"Sec-fetch-mode":     "cors",
	"Sec-fetch-site":     "same-site",
	"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0",
}

func NewRealtimeQuoteRequest() *RealtimeQuoteRequest {
	return &RealtimeQuoteRequest{}
}

func (r *RealtimeQuoteRequest) SetHeader(key, value string) {
	r.Request.Header.Set(key, value)
}

func (r *RealtimeQuoteRequest) BuildRequest() {

}
