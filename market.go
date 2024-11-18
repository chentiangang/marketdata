package dongfang

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chentiangang/xlog"
)

type Request struct {
	UrlValues url.Values
	Host      string
	Market    string
}

type Resp struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int64  `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts,omitempty"`
	Data   struct {
		Total int `json:"total"`
		Diff  []struct {
			//F1   int     `json:"f1,omitempty"`
			F2 interface{} `json:"f2,omitempty"`
			F3 interface{} `json:"f3,omitempty"` // 涨跌幅
			F4 interface{} `json:"f4,omitempty"`
			//F5   int     `json:"f5,omitempty"`
			//F6   float64 `json:"f6,omitempty"`
			//F7   int     `json:"f7,omitempty"`
			F8 interface{} `json:"f8,omitempty"`
			//F9   int     `json:"f9,omitempty"`
			//F10  int     `json:"f10,omitempty"`
			F12 string `json:"f12,omitempty"`
			F13 int    `json:"f13,omitempty"` // 交易所 0上交所 1深交所
			F14 string `json:"f14,omitempty"`
			//F18  int     `json:"f18,omitempty"`
			//F19  int     `json:"f19,omitempty"`
			F20 interface{} `json:"f20,omitempty"`
			F21 interface{} `json:"f21,omitempty"`
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
	} `json:"data,omitempty"`
}

const listApi = "/api/qt/clist/get"

func (req *Request) Set(key, value string) {
	req.UrlValues.Set(key, value)
}

func (req *Request) url(pageNum, pageSiz int) (string, error) {
	req.Host = Domain()

	url, err := url.Parse(fmt.Sprintf("%s%s", req.Host, listApi))
	if err != nil {
		return "", err
	}
	req.UrlValues = url.Query()
	req.Set("cb", "jQuery11240699042934591428_1726233885825")

	req.Set("pn", fmt.Sprintf("%d", pageNum))
	req.Set("pz", fmt.Sprintf("%d", pageSiz))
	req.Set("po", "0")
	req.Set("np", "1")
	req.Set("ut", UT())
	req.Set("fltt", "2")
	req.Set("invt", "2")
	req.Set("dect", "1")
	//req.Set("wbp2u")
	req.Set("fid", "f12")
	//req.Set("fs", "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048")
	req.Set("fs", req.Market)

	req.Set("fields", "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152")
	req.Set("-", fmt.Sprintf("%d", time.Now().UnixMilli()))
	url.RawQuery = req.UrlValues.Encode()
	return url.String(), nil
}

const (
	ChinaMarket = "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23"
	UsMarket    = "b:MK0001"
)

func (r *Request) Request(pageNum, pageSize int) (*http.Response, error) {
	client := &http.Client{}
	u, err := r.url(pageNum, pageSize)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("accept-encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Content-type", "text/event-stream")
	req.Header.Set("Accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", r.Host)
	req.Header.Set("Origin", "https://quote.eastmoney.com")
	req.Header.Set("Referer", "https://quote.eastmoney.com/zixuan/?from=home")
	req.Header.Set("Sec-ch-ua", `"Chromium";v="128", "Not;A=Brand";v="24", "Microsoft Edge";v="128"`)
	req.Header.Set("Sec-ch-ua-mobile", "?0")
	req.Header.Set("Sec-ch-ua-platform", "macOS")
	req.Header.Set("Sec-fetch-dest", "empty")
	req.Header.Set("Sec-fetch-mode", "cors")
	req.Header.Set("Sec-fetch-site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0")
	return client.Do(req)
}

func (r *Request) Parse(respon *http.Response) (qs []quote.Stock, err error) {

	bs, err := Unzip(respon)
	defer respon.Body.Close()
	if err != nil {
		xlog.Error("reader error: %s", err)
		return nil, nil
	}
	trim := strings.Trim(string(bs), "jQuery11240699042934591428_1726233885825(")
	trim = strings.Trim(trim, ");")
	var resp Resp

	//fmt.Println(trim)
	err = json.Unmarshal([]byte(trim), &resp)
	if err != nil {
		xlog.Error("Error decoding JSON: %s", err)
		return
	}

	if len(resp.Data.Diff) == 0 {
		return nil, errors.New("no change")
	}

	for _, v := range resp.Data.Diff {
		// 过滤退市
		if fmt.Sprintf("%v", v.F2) == "-" {
			continue
		}
		var q quote.Stock
		q.Name = v.F14
		q.Symbol = v.F12
		q.Price = math.ConvertToFloat64(v.F2)
		q.PriceLimit = math.ConvertToFloat64(v.F3)
		q.DifferenceValue = math.ConvertToFloat64(v.F4)
		q.TurnoverRate = math.ConvertToFloat64(v.F8)
		q.Exchange = v.F13
		q.Alias = alias(strings.TrimSpace(v.F14))
		q.TotalValue = math.ConvertToInt(v.F20)
		q.CirculatingValue = math.ConvertToInt(v.F21)
		qs = append(qs, q)
	}
	return qs, nil
}

func alias(name string) string {
	a := pinyin.NewArgs()
	a.Heteronym = true
	if name == "" {
		return name
	}

	names := strings.Split(name, "")

	var s1 string
	for _, n := range names {
		for _, i := range pinyin.Pinyin(n, a) {
			if exist(n) {
				s1 = s1 + i[1][:1]
				break
			}
			s1 = s1 + i[0][:1]

		}
	}

	var s2 string
	for _, n := range names {
		for _, i := range pinyin.Pinyin(n, a) {
			if exist(n) {
				s2 = s2 + i[1]
				break
			}
			s2 = s2 + i[0]

		}
	}

	return s1 + "-" + s2
}

func exist(name string) bool {
	s := strings.Split(name, "")
	for _, v := range s {
		switch v {
		case "行":
			return true
		}
	}
	return false
}
