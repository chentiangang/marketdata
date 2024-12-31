package dongfang

import (
	"time"

	"github.com/cinar/indicator/v2/asset"
)

// 合并一组K线数据
func mergePeriod(klines []*asset.Snapshot) *asset.Snapshot {
	merged := klines[0] // 以第一根K线为基础
	for _, k := range klines[1:] {
		merged.High = max(merged.High, k.High)
		merged.Low = min(merged.Low, k.Low)
		merged.Close = k.Close
		merged.Volume += k.Volume
	}

	return merged
}

// 获取较大的数值
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 获取较小的数值
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

//
//// 合并K线到45分钟的规则
//func Merge5MinTo45MinKlines(klines []*asset.Snapshot) []*asset.Snapshot {
//	var result []*asset.Snapshot
//	i := 0
//
//	for i < len(klines) {
//		k := klines[i]
//		hour, minute := k.Date.Hour(), k.Date.Minute()
//
//		switch {
//		case hour == 9 && minute == 45:
//			// 合并 9:45, 10:00, 10:15 为 9:30, 如果不足3根, 合并所有剩余的K线
//			if i+2 < len(klines) {
//				tempKline := mergePeriod(klines[i : i+3])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 9, 30, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i += 3
//			} else {
//				tempKline := mergePeriod(klines[i:])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 9, 30, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i = len(klines) // 处理完最后一组K线，退出循环
//			}
//
//		case hour == 10 && minute == 30:
//			// 合并 10:30, 10:45, 11:00 为 10:15, 如果不足3根, 合并所有剩余的K线
//			if i+2 < len(klines) {
//				tempKline := mergePeriod(klines[i : i+3])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 10, 15, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i += 3
//			} else {
//				tempKline := mergePeriod(klines[i:])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 10, 15, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i = len(klines) // 处理完最后一组K线，退出循环
//			}
//
//		case hour == 11 && minute == 15:
//			// 合并 11:15, 11:30 为 11:00, 如果不足2根, 合并所有剩余的K线
//			if i+1 < len(klines) {
//				tempKline := mergePeriod(klines[i : i+2])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 11, 0, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i += 2
//			} else {
//				tempKline := mergePeriod(klines[i:])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 11, 0, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i = len(klines) // 处理完最后一组K线，退出循环
//			}
//
//		case hour == 13 && minute == 15:
//			// 合并单根K线 (13:15 -> 12:30)
//			tempKline := klines[i]
//			tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 12, 30, 0, 0, time.Local)
//			result = append(result, tempKline)
//			i++
//
//		case hour == 13 && minute == 30:
//			// 合并 13:30, 13:45, 14:00 为 13:15, 如果不足3根, 合并所有剩余的K线
//			if i+2 < len(klines) {
//				tempKline := mergePeriod(klines[i : i+3])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 13, 15, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i += 3
//			} else {
//				tempKline := mergePeriod(klines[i:])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 13, 15, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i = len(klines) // 处理完最后一组K线，退出循环
//			}
//
//		case hour == 14 && minute == 15:
//			// 合并 14:15, 14:30, 14:45 为 14:00, 如果不足3根, 合并所有剩余的K线
//			if i+2 < len(klines) {
//				tempKline := mergePeriod(klines[i : i+3])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 14, 0, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i += 3
//			} else {
//				tempKline := mergePeriod(klines[i:])
//				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 14, 0, 0, 0, time.Local)
//				result = append(result, tempKline)
//				i = len(klines) // 处理完最后一组K线，退出循环
//			}
//
//		case hour == 15 && minute == 0:
//			// 合并 15:00 为 14:45, 仅在处理当天的最后一根
//			tempKline := klines[i]
//			tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 14, 45, 0, 0, time.Local)
//			result = append(result, tempKline)
//			i++ // 继续处理下一天的数据
//
//		default:
//			i++
//		}
//	}
//
//	return result
//}

// 合并K线到45分钟的规则
func MergeTo45MinKlines(klines []*asset.Snapshot) []*asset.Snapshot {
	var result []*asset.Snapshot
	i := 0

	for i < len(klines) {
		k := klines[i]
		hour, minute := k.Date.Hour(), k.Date.Minute()

		switch {
		case hour == 9 && minute == 45:
			// 合并 9:45, 10:00, 10:15 为 9:30, 如果不足3根, 合并所有剩余的K线
			if i+2 < len(klines) {
				tempKline := mergePeriod(klines[i : i+3])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 9, 30, 0, 0, time.Local)
				result = append(result, tempKline)
				i += 3
			} else {
				tempKline := mergePeriod(klines[i:])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 9, 30, 0, 0, time.Local)
				result = append(result, tempKline)
				i = len(klines) // 处理完最后一组K线，退出循环
			}

		case hour == 10 && minute == 30:
			// 合并 10:30, 10:45, 11:00 为 10:15, 如果不足3根, 合并所有剩余的K线
			if i+2 < len(klines) {
				tempKline := mergePeriod(klines[i : i+3])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 10, 15, 0, 0, time.Local)
				result = append(result, tempKline)
				i += 3
			} else {
				tempKline := mergePeriod(klines[i:])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 10, 15, 0, 0, time.Local)
				result = append(result, tempKline)
				i = len(klines) // 处理完最后一组K线，退出循环
			}

		case hour == 11 && minute == 15:
			// 合并 11:15, 11:30 为 11:00, 如果不足2根, 合并所有剩余的K线
			if i+1 < len(klines) {
				tempKline := mergePeriod(klines[i : i+2])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 11, 0, 0, 0, time.Local)
				result = append(result, tempKline)
				i += 2
			} else {
				tempKline := mergePeriod(klines[i:])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 11, 0, 0, 0, time.Local)
				result = append(result, tempKline)
				i = len(klines) // 处理完最后一组K线，退出循环
			}

		case hour == 13 && minute == 15:
			// 合并单根K线 (13:15 -> 12:30)
			tempKline := klines[i]
			tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 12, 30, 0, 0, time.Local)
			result = append(result, tempKline)
			i++

		case hour == 13 && minute == 30:
			// 合并 13:30, 13:45, 14:00 为 13:15, 如果不足3根, 合并所有剩余的K线
			if i+2 < len(klines) {
				tempKline := mergePeriod(klines[i : i+3])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 13, 15, 0, 0, time.Local)
				result = append(result, tempKline)
				i += 3
			} else {
				tempKline := mergePeriod(klines[i:])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 13, 15, 0, 0, time.Local)
				result = append(result, tempKline)
				i = len(klines) // 处理完最后一组K线，退出循环
			}

		case hour == 14 && minute == 15:
			// 合并 14:15, 14:30, 14:45 为 14:00, 如果不足3根, 合并所有剩余的K线
			if i+2 < len(klines) {
				tempKline := mergePeriod(klines[i : i+3])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 14, 0, 0, 0, time.Local)
				result = append(result, tempKline)
				i += 3
			} else {
				tempKline := mergePeriod(klines[i:])
				tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 14, 0, 0, 0, time.Local)
				result = append(result, tempKline)
				i = len(klines) // 处理完最后一组K线，退出循环
			}

		case hour == 15 && minute == 0:
			// 合并 15:00 为 14:45, 仅在处理当天的最后一根
			tempKline := klines[i]
			tempKline.Date = time.Date(k.Date.Year(), k.Date.Month(), k.Date.Day(), 14, 45, 0, 0, time.Local)
			result = append(result, tempKline)
			i++ // 继续处理下一天的数据

		default:
			i++
		}
	}

	return result
}
