package dongfang

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/chentiangang/marketdata/model"
	"github.com/chentiangang/marketdata/util"

	"github.com/chentiangang/xlog"
	"github.com/duke-git/lancet/v2/strutil"
)

type Quote struct {
	BaseURL string
	Symbols []string
	Request *http.Request
	ctx     context.Context
	cancel  context.CancelFunc
}

type QuoteResponse struct {
	Rc     int        `json:"rc"`
	Rt     int        `json:"rt"`
	Svr    int64      `json:"svr"`
	Lt     int        `json:"lt"`
	Full   int        `json:"full"`
	Dlmkts string     `json:"dlmkts,omitempty"`
	Data   *QuoteData `json:"data,omitempty"`
}

type QuoteData struct {
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
	"Content-type":       "text/event-stream",
	"Accept-language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	"Cache-Control":      "no-cache",
	"Connection":         "keep-alive",
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

func NewRealtimeQuoteRequest() *Quote {
	ctx, cancel := context.WithCancel(context.Background())
	return &Quote{
		BaseURL: HttpHost() + "/api/qt/ulist/sse",
		cancel:  cancel,
		ctx:     ctx,
	}
}

func (r *Quote) SetHeader(key, value string) {
	r.Request.Header.Set(key, value)
}

func (r *Quote) BuildRequest() error {
	u, err := url.Parse(r.BaseURL)
	if err != nil {
		xlog.Error("%s", err)
		return err
	}
	query := u.Query()
	query.Set("secids", strings.Join(r.Symbols, ","))
	query.Set("fields", "f12,f13,f19,f14,f139,f148,f2,f4,f1,f125,f18,f3,f152,f5,f30,f31,f32,f6,f8,f7,f10,f22,f9,f112,f100")
	query.Set("invt", "3")
	query.Set("ut", ut())
	query.Set("fid", "")
	query.Set("mpi", "1000")
	query.Set("po", "1")
	query.Set("pi", "0")
	query.Set("pz", fmt.Sprintf("%d", len(r.Symbols)))
	query.Set("dect", "1")
	u.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		xlog.Error("%s", err)
		return err
	}
	r.Request = req

	// Set headers
	for k, v := range defaultRealTimeQuoteHeaders {
		r.SetHeader(k, v)
	}
	return nil
}

func (r *Quote) Set(symbols []string) {
	r.Symbols = symbols
}

func (r *Quote) Fetch() chan []model.QuotePtr {
	ch := make(chan []model.QuotePtr)

	go func() {
		err := r.BuildRequest()
		if err != nil {
			xlog.Error("%s", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(r.Request)
		defer close(ch)
		defer resp.Body.Close()
		if err != nil {
			xlog.Error("%s", err)
			return
		}

		if err != nil {
			xlog.Error("%s", err)
			return
		}
		scanner := bufio.NewScanner(resp.Body)
		for {
			select {
			case <-r.ctx.Done():
				return
			default:
				if !scanner.Scan() {
					if err := scanner.Err(); err != nil {
						xlog.Error("SseSub scanner error: %s", err)
					}
					return
				}
				line := scanner.Text()
				//fmt.Println(line)
				if len(line) > 0 {
					data, err := r.parse([]byte(line))
					//fmt.Println(data)
					if err != nil {
						xlog.Error("%s", err)
						continue
					}
					fmt.Println("write", len(data))
					if len(data) > 0 {
						ch <- data
					}
				}
			}
		}
	}()
	return ch
}

func (r *Quote) parse(bs []byte) (qs []model.QuotePtr, err error) {
	trim := strutil.Trim(string(bs), "data:")
	var resp QuoteResponse
	err = json.Unmarshal([]byte(trim), &resp)
	if err != nil {
		xlog.Error("%s", err)
		xlog.Error("%s", string(bs))
		return qs, err
	}

	if resp.Data == nil {
		return nil, err
	}

	if len(resp.Data.Diff) == 0 {
		return nil, errors.New("no change")
	}

	for _, v := range resp.Data.Diff {
		var q model.QuotePtr
		q.Name = v.F14
		q.Symbol = v.F12

		if v.F2 != nil {
			q.Price = new(float64)
			*q.Price = util.DivideByHundred(*v.F2)
		}

		if v.F3 != nil {
			q.PriceLimit = new(float64)
			*q.PriceLimit = util.DivideByHundred(*v.F3)
		}

		if v.F4 != nil {
			q.DifferenceValue = new(float64)
			*q.DifferenceValue = util.DivideByHundred(*v.F4)
		}

		if v.F8 != nil {
			q.TurnoverRate = new(float64)
			*q.TurnoverRate = util.DivideByHundred(*v.F8)
		}

		if v.F20 != nil {
			q.TotalValue = new(int64)
			*q.TotalValue = util.ConvertToInt(v.F20)
		}

		if v.F21 != nil {
			q.CirculatingValue = new(int64)
			*q.CirculatingValue = util.ConvertToInt(v.F21)
		}

		if v.F84 != nil {
			q.TotalShares = new(int64)
			*q.TotalShares = util.ConvertToInt(v.F84)
		}

		qs = append(qs, q)
	}
	return qs, nil
}

func (r *Quote) Close() {
	r.cancel()
}
