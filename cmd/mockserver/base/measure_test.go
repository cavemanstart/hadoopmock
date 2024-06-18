package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	chargeMap = map[int64]int64{
		100: 1,
		101: 2,
		102: 3,
		103: 4,
		104: 5,
		105: 6,
		106: 7,
		107: 8,
		108: 9,
		109: 10,
	}
	chargeMa2 = map[int64]int64{
		100: 11,
		101: 12,
		102: 13,
		103: 14,
		104: 15,
		105: 16,
		106: 17,
		107: 18,
		108: 19,
		109: 20,
	}
)

func TestDay95Measure(t *testing.T) {
	measure, _ := Day95Measure(chargeMap)
	ast := assert.New(t)
	ast.Equal(measure, int64(10))
}

func TestDayPeakMeasure(t *testing.T) {
	measure, _ := DayPeakMeasure(chargeMap)
	ast := assert.New(t)
	ast.Equal(measure, int64(10))
}
func TestIntervalPeak(t *testing.T) {
	list := []map[int64]int64{chargeMap, chargeMa2}
	peak, _ := IntervalPeak(list)
	ast := assert.New(t)
	ast.Equal(peak, int64(19))
}
func TestIntervalBillPeak(t *testing.T) {
	list := []map[int64]int64{chargeMap, chargeMa2}
	biilPeak, _ := IntervalBillPeak(list, 3)
	ast := assert.New(t)
	ast.Equal(biilPeak, int64(19))
}
