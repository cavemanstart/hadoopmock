package mock

import (
	"github.com/zeromicro/go-zero/core/logx"
	"hadoopmock/cmd/internal/common"
	"math/rand"
	"strconv"
	"time"
)

func MockVendorEveningMetric(startDate string, endDate string) *common.MeasureCommonData {
	if startDate > endDate {
		logx.Errorf("invalid input: %v > %v", startDate, endDate)
		return nil
	}
	day := startDate
	dayPeak := map[string]common.MeasureCommonUnit{}
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
		dayPeak[day] = common.MeasureCommonUnit{
			Bandwidth: tmpBandwidth,
			Time:      d.Unix() + 64800 + rand.Int63n(21600), //晚高峰时段
		}
		d = d.AddDate(0, 0, 1)
		day = d.Format("2006-01-02")
		if day == endDay.AddDate(0, 0, 1).Format("2006-01-02") {
			break
		}
		cnt++
	}
	mockMeasureCommonData := common.MeasureCommonData{
		Id:      idPrefix + strconv.FormatInt(time.Now().UnixMilli()+rand.Int63n(10000), 10),
		DayPeak: dayPeak,
	}
	return &mockMeasureCommonData
}
