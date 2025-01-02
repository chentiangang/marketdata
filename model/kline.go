package model

import (
	"github.com/cinar/indicator/v2/asset"
)

type Kline struct {
	Name      string
	Symbol    string
	Snapshots []*asset.Snapshot
}

// Last 用来获取最后第n根K线的数据
func (k *Kline) Last(n int) *asset.Snapshot {
	if len(k.Snapshots) < n {
		return nil
	}
	return k.Snapshots[len(k.Snapshots)-n]
}

// First 用来获取从最开始的第n根K线
func (k *Kline) First(n int) *asset.Snapshot {
	if len(k.Snapshots) < n {
		return nil
	}
	return k.Snapshots[n]
}

// LastPrices 返回最后n根K线的价格
func (k *Kline) LastPrices(n int) []float64 {
	// 如果K线数量小于n，则返回所有K线的价格
	if len(k.Snapshots) < n {
		n = len(k.Snapshots)
	}

	// 创建一个切片来存储最后n根K线的价格
	var prices []float64
	for i := len(k.Snapshots) - n; i < len(k.Snapshots); i++ {
		// 假设Snapshot有一个Close字段表示K线价格
		prices = append(prices, k.Snapshots[i].Close)
	}

	return prices
}

// LastLows 返回最后n根K线中的最低价
func (k *Kline) LastLows(n int) []float64 {
	// 如果K线数量小于n，则返回所有K线的价格
	if len(k.Snapshots) < n {
		n = len(k.Snapshots)
	}

	// 创建一个切片来存储最后n根K线的价格
	var prices []float64
	for i := len(k.Snapshots) - n; i < len(k.Snapshots); i++ {
		// 假设Snapshot有一个Close字段表示K线价格
		prices = append(prices, k.Snapshots[i].Low)
	}

	return prices
}
