package service

import (
	"github.com/zeromicro/go-zero/core/logx"
	"hadoopmock/cmd/internal/types"
	"math/rand"
	"time"
)

func MockCustomerPerNodeMetric(startDate string, endDate string, days int64, nodeIdslice []string) *types.MeasureCommonDataNodes {
	if startDate > endDate {
		logx.Errorf("invalid input: %v > %v", startDate, endDate)
		return nil
	}
	endDay, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
	if err != nil {
		logx.Errorf("parse api result err: %v", err)
	}
	slice := []*types.MeasureCommonDataPerNode{}
	for _, nodeId := range nodeIdslice {
		var cnt int64 = 0
		curNodeData := types.MeasureCommonDataPerNode{}
		curNodeData.NodeId = nodeId
		day := startDate
		dayPeak := map[string]*types.MeasureCommonUnit{}
		peak := types.MeasureCommonUnit{}
		billPeak := types.MeasureCommonUnit{}
		sumBandwidth := int64(0)
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
		curNodeData.DayPeak = dayPeak
		curNodeData.Peak = billPeak
		curNodeData.BillPeak = billPeak
		slice = append(slice, &curNodeData)
	}
	res := &types.MeasureCommonDataNodes{
		NodeList: slice,
	}
	return res
}
