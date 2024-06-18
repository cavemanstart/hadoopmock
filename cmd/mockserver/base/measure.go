package base

import (
	"sort"
)

func DayPeakMeasure(flow map[int64]int64) (int64, int64) {
	var dayPeak int64 = 0
	var timePeak int64 = 0
	for key, value := range flow {
		if value > dayPeak {
			timePeak = key
			dayPeak = value
		}
	}
	return dayPeak, timePeak
}
func Day95Measure(flow map[int64]int64) (int64, int64) {
	length := len(flow)
	var k = int(float64(length) * 0.05)
	values := []int64{}
	for _, val := range flow {
		values = append(values, val)
	}
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	for key, val := range flow {
		if val == values[k] {
			return val, key
		}
	}
	return 0, 0
}
func IntervalPeak(mapList []map[int64]int64) (int64, int64) {
	values := []int64{}
	for _, mp := range mapList {
		for _, value := range mp {
			values = append(values, value)
		}
	}
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	k := int(float64(len(values)) * 0.05)
	for _, mp := range mapList {
		for key, value := range mp {
			if value == values[k] {
				return value, key
			}
		}
	}
	return 0, 0
}
func IntervalBillPeak(mapList []map[int64]int64, billDays int64) (int64, int64) {
	values := []int64{}
	for _, mp := range mapList {
		for _, value := range mp {
			values = append(values, value)
		}
	}
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	k := int(float64(len(values)) * float64(billDays) / float64(len(mapList)) * 0.05)
	for _, mp := range mapList {
		for key, value := range mp {
			if value == values[k] {
				return value, key
			}
		}
	}
	return 0, 0
}
