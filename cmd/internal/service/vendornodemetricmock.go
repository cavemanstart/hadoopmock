package service

import (
	"math/rand"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"hadoopmock/cmd/internal/types"
)

var (
	baseBandwidth  int64  = 20000000000
	rangeBandwidth int64  = 10000000000
	idPrefix       string = "test-"
)

func MockVendorNodeMetric(startDate string, endDate string, days int64) *types.MeasureCommonData {
	if startDate > endDate {
		logx.Errorf("invalid input: %v > %v", startDate, endDate)
		return nil
	}
	day := startDate
	dayPeak := map[string]*types.MeasureCommonUnit{}
	peak := types.MeasureCommonUnit{}
	billPeak := types.MeasureCommonUnit{}
	sumBandwidth := int64(0)
	endDay, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
	if err != nil {
		logx.Errorf("parse api result err: %v", err)
	}
	var cnt int64 = 0
	for {
		d, err := time.ParseInLocation("2006-01-02", day, time.Local)
		if err != nil {
			logx.Errorf("parse api result err: %v", err)
		}
		tmpBandwidth := baseBandwidth + rand.Int63n(rangeBandwidth)
		sumBandwidth += tmpBandwidth
		dayPeak[day] = &types.MeasureCommonUnit{
			Bandwidth: tmpBandwidth,
			Time:      d.Unix() + rand.Int63n(86400), //一天有86400秒
		}
		d = d.AddDate(0, 0, 1)
		day = d.Format("2006-01-02")
		if day == endDay.AddDate(0, 0, 1).Format("2006-01-02") {
			break
		}
		cnt++
	}
	peak.Bandwidth = sumBandwidth / cnt
	peak.Time = dayPeak[startDate].Time
	billPeak.Bandwidth = sumBandwidth / days
	billPeak.Time = peak.Time
	mockMeasureCommonData := types.MeasureCommonData{
		DayPeak:  dayPeak,
		Peak:     peak,
		BillPeak: billPeak,
	}
	return &mockMeasureCommonData
}
