package dongfang

import (
	"encoding/json"
	"fmt"
	"marketdata/model"
	"marketdata/util"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chentiangang/xlog"
	"github.com/cinar/indicator/v2/asset"
)

type KlineResponse struct {
	Data   KlineData `json:"data"`
	Dlmkts string    `json:"dlmkts"`
	Full   int       `json:"full"`
	Lt     int       `json:"lt"`
	Rc     int       `json:"rc"`
	Rt     int       `json:"rt"`
	Svr    int       `json:"svr"`
}

type KlineData struct {
	Code      string   `json:"code"`
	Decimal   int      `json:"decimal"`
	Dktotal   int      `json:"dktotal"`
	Klines    []string `json:"klines"`
	Market    int      `json:"market"`
	Name      string   `json:"name"`
	PreKPrice float64  `json:"preKPrice"`
}

type KlineRequest struct {
	BaseURL string
	Request *http.Request
}

var defaultKlineHeaders = map[string]string{
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"Accept-Encoding":           "gzip, deflate, br, zstd",
	"Accept-Language":           "zh-CN,zh;q=0.9",
	"Cache-Control":             "max-age=0",
	"Connection":                "keep-alive",
	"Host":                      "push2his.eastmoney.com",
	"Sec-CH-UA":                 `"Chromium";v="128", "Not;A=Brand";v="24", "Microsoft Edge";v="128"`,
	"Sec-CH-UA-Mobile":          "?0",
	"Sec-CH-UA-Platform":        `"Windows"`,
	"Sec-Fetch-Dest":            "document",
	"Sec-Fetch-Mode":            "navigate",
	"Sec-Fetch-Site":            "none",
	"Sec-Fetch-User":            "?1",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0",
}

func NewDefaultKlineRequest() *KlineRequest {
	return &KlineRequest{
		BaseURL: "https://push2his.eastmoney.com" + "/api/qt/stock/kline/get",
	}
}

// SetHeader adds or updates a header in the request.
func (kr *KlineRequest) SetHeader(key, value string) {
	kr.Request.Header.Set(key, value)
}

func (kr *KlineRequest) BuildRequest(symbol, period, limit string) error {
	u, err := url.Parse(kr.BaseURL)
	if err != nil {
		xlog.Error("%s")
		return err
	}
	query := u.Query()
	query.Set("cb", "jQuery351029107463534780975_1726757437952")
	query.Set("secid", symbol)
	query.Set("ut", ut())
	query.Set("fields1", "f1,f2,f3,f4,f5,f6")
	query.Set("fields2", "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61")
	query.Set("klt", period)
	query.Set("fqt", "1")
	query.Set("end", "20500101")
	query.Set("lmt", limit)
	query.Set("_", fmt.Sprintf("%d", time.Now().UnixMilli()))
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		xlog.Error("%s", err)
		return err
	}
	kr.Request = req

	// Set headers
	for k, v := range defaultKlineHeaders {
		kr.SetHeader(k, v)
	}
	return nil
}

func (kr *KlineRequest) Fetch(symbol, period, limit string) (model.Kline, error) {
	err := kr.BuildRequest(symbol, period, limit)
	if err != nil {
		xlog.Error("%s", err)
		return model.Kline{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(kr.Request)
	if err != nil {
		xlog.Error("%s", err)
		return model.Kline{}, err
	}

	bs, err := Unzip(resp)
	defer resp.Body.Close()
	if err != nil {
		xlog.Error("reader error: %s", err)
		return model.Kline{}, err
	}
	trim := strings.Trim(string(bs), "jQuery351029107463534780975_1726757437952(")
	trim = strings.Trim(trim, ");")

	var klineResp KlineResponse
	err = json.Unmarshal([]byte(trim), &klineResp)
	if err != nil {
		xlog.Error("Error decoding JSON: %s", err)
		return model.Kline{}, err
	}
	var kline model.Kline
	kline.Symbol = klineResp.Data.Code
	kline.Name = klineResp.Data.Name
	for _, i := range klineResp.Data.Klines {
		k := &asset.Snapshot{}
		row := strings.Split(i, ",")
		k.Date, _ = time.Parse("2006-01-02 15:04", row[0])
		k.Open = util.ConvertToFloat64(row[1])   // 开盘价
		k.Close = util.ConvertToFloat64(row[2])  // 收盘价
		k.High = util.ConvertToFloat64(row[3])   // 最高价
		k.Low = util.ConvertToFloat64(row[4])    // 最低价
		k.Volume = util.ConvertToFloat64(row[5]) // 成交量
		kline.Snapshots = append(kline.Snapshots, k)
	}
	return kline, nil
}
