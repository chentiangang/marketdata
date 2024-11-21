package util

import (
	"fmt"
	"strconv"

	"github.com/chentiangang/xlog"
)

func DivideByHundred(i int) float64 {
	var f float64

	f = float64(i) / float64(100)
	return f
}

// ConvertToInt 尝试将 interface{} 转换为 int 类型
func ConvertToInt(i interface{}) int64 {
	switch v := i.(type) {
	case float64:
		return int64(v)
	case int:
		return int64(v)
	case int64:
		return v
	case string:
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			xlog.Error("%s", err)
			return 0
		}
		return n
	default:
		return 0
	}
}

// ConvertToFloat64 尝试将 interface{} 转换为 float64 类型
func ConvertToFloat64(i interface{}) float64 {
	switch v := i.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			xlog.Error("%s", err)
			return 0
		}
		return f
	default:
		return 0
	}
}

// ConvertToFormattedUnit 将 int 转换为 float64，并根据数值大小转换为 "万"、"亿" 或 "万亿"
func ConvertToFormattedUnit(value int64) string {
	var result float64
	var unit string

	switch {
	case value >= 1000000000000: // 如果值大于等于1万亿
		result = float64(value) / 1000000000000
		unit = "万亿"
	case value >= 100000000: // 如果值大于等于1亿
		result = float64(value) / 100000000
		unit = "亿"
	case value >= 10000: // 如果值大于等于1万
		result = float64(value) / 10000
		unit = "万"
	default: // 如果值小于1万，不需要单位
		result = float64(value)
		unit = ""
	}

	// 格式化结果并返回带单位的字符串，保留两位小数
	return fmt.Sprintf("%.2f%s", result, unit)
}

// ConvertToLargeUnit 将流通市值转换为 "万亿" 或 "亿" 单位的浮点数
//func ConvertToLargeUnit(value int64) string {
//	var result float64
//	var unit string
//
//	if value >= 1000000000000 { // 如果值大于等于1万亿
//		result = float64(value) / 1000000000000 // 转换为万亿
//		unit = "万亿"
//	} else if value >= 100000000 { // 如果值大于等于1亿
//		result = float64(value) / 100000000 // 转换为亿
//		unit = "亿log.Println("Error reading from stream:", err)"
//	} else {
//		result = float64(value)
//		unit = "" // 如果值小于1亿，不加单位
//	}
//
//	// 格式化结果并返回带单位的字符串，保留两位小数
//	return fmt.Sprintf("%.2f%s", result, unit)
//}

// ConvertToLargeUnit 将流通市值转换为 "万亿"、"亿" 或 "万" 单位的浮点数
func ConvertToLargeUnit(value int64) string {
	var result float64
	var unit string

	switch {
	case value >= 1000000000000: // 如果值大于等于1万亿
		result = float64(value) / 1000000000000
		unit = "万亿"
	case value >= 100000000: // 如果值大于等于1亿
		result = float64(value) / 100000000
		unit = "亿"
	case value >= 10000: // 如果值大于等于1万
		result = float64(value) / 10000
		unit = "万"
	default: // 如果值小于1万，不加单位
		result = float64(value)
		unit = ""
	}

	// 格式化结果并返回带单位的字符串，保留两位小数
	return fmt.Sprintf("%.2f%s", result, unit)
}
