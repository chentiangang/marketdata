package model

import (
	"github.com/cinar/indicator/v2/asset"
)

type Kline struct {
	Name      string
	Symbol    string
	Snapshots []asset.Snapshot
}
