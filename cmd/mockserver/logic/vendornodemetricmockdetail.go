package logic

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"hadoopmock/cmd/mockserver/base"
	"hadoopmock/cmd/mockserver/config"
	"hadoopmock/cmd/mockserver/types"
)

type VendorNodeMetricDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewVendorNodeMetricDetailLogic(ctx context.Context) *VendorNodeMetricDetailLogic {
	return &VendorNodeMetricDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *VendorNodeMetricDetailLogic) VendorNodeMetricMockDetail(sc *config.ServiceContext, req *types.PostVendorNodeMetricReq) (*types.MeasureCommonData, error) {
	startDate := req.Start
	endDate := req.End
	if startDate > endDate {
		logx.Errorf("invalid args: %v > %v", startDate, endDate)
		return nil, errors.New("invalid args")
	}
	resp := &types.MeasureCommonData{} //初始化
	m := sc.JarvisChargeNode5minModel
	nodeIds := []string{}
	for _, measureMetricNode := range req.Nodes {
		nodeIds = append(nodeIds, measureMetricNode.NodeId)
	}

	measureType := req.MeasureType
	day := startDate
	endDay, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
	if err != nil {
		logx.Errorf("parse api result err: %v", err)
	}
	dayPeak := map[string]*types.MeasureCommonUnit{}
	var cntDay int64 = 0
	intervalMeasure := []map[int64]int64{}
	for day != endDay.AddDate(0, 0, 1).Format("2006-01-02") {
		d, err := time.ParseInLocation("2006-01-02", day, time.Local)
		if err != nil {
			logx.Errorf("parse api result err: %v", err)
			return nil, err
		}

		measureList := [][]int64{}
		for _, nodeId := range nodeIds {
			data, err := m.FindByNodeIdAndDate(sc.Context, nodeId, d)
			if err != nil {
				logx.Errorf("db query error: %v", err)
				return nil, err
			}
			measureList = append(measureList, data.OutFlow)
		}
		totalMeasure := make(map[int64]int64)
		for _, measure := range measureList {
			for i, charge := range measure {
				totalMeasure[d.Unix()+int64(i*300)] += charge
			}
		}
		intervalMeasure = append(intervalMeasure, totalMeasure)
		var peak int64 = 0
		var t int64 = 0
		if measureType == "peak" {
			peak, t = base.DayPeakMeasure(totalMeasure)
		} else if measureType == "dayPeak" {
			peak, t = base.Day95Measure(totalMeasure)
		}
		dayPeak[day] = &types.MeasureCommonUnit{peak, t}
		d = d.AddDate(0, 0, 1)
		day = d.Format("2006-01-02")
		cntDay++
	}
	resp.DayPeak = dayPeak
	peak, ts := base.IntervalPeak(intervalMeasure)
	resp.Peak = types.MeasureCommonUnit{peak, ts}
	billPeak, ts2 := base.IntervalBillPeak(intervalMeasure, req.BillDays)
	resp.BillPeak = types.MeasureCommonUnit{billPeak, ts2}
	return resp, nil
}
