package service

import (
	"github.com/zeromicro/go-zero/core/logx"
	"hadoopmock/cmd/internal/types"
	"math/rand"
	"time"
)

func MockPingLossNodeRate(nodeIds []string, startDate string, endDate string) *types.PingLossNodeRatioData {
	if startDate > endDate {
		logx.Errorf("invalid input: %v > %v", startDate, endDate)
		return nil
	}
	endDay, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
	if err != nil {
		logx.Errorf("parse api result err: %v", err)
	}
	pingLossRatio := map[string][]*types.NodePingLossRatioUnit{}
	day := startDate
	for {
		d, err := time.ParseInLocation("2006-01-02", day, time.Local)
		if err != nil {
			logx.Errorf("parse api result err: %v", err)
		}
		slice := []*types.NodePingLossRatioUnit{}
		for _, nodeId := range nodeIds {
			unit := types.NodePingLossRatioUnit{
				NodeId: nodeId,
				Ratio:  rand.Float64() / 2,
			}
			slice = append(slice, &unit)
		}
		pingLossRatio[day] = slice
		d = d.AddDate(0, 0, 1)
		day = d.Format("2006-01-02")
		if day == endDay.AddDate(0, 0, 1).Format("2006-01-02") {
			break
		}
	}
	return &types.PingLossNodeRatioData{
		PingLossRatio: pingLossRatio,
	}
}
