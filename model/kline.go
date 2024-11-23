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
