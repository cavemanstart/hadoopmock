package service

import (
	"hadoopmock/cmd/internal/types"
	"math/rand"
)

func MockCustomerNode5Min(startTime int64, endTime int64) *types.MeasureCommonUnitList {
	commonUnitList := []*types.MeasureCommonUnit{}
	timestamp := startTime
	for {
		if timestamp >= endTime {
			break
		}
		unit := types.MeasureCommonUnit{
			Bandwidth: baseBandwidth + rand.Int63n(rangeBandwidth),
			Time:      timestamp,
		}
		commonUnitList = append(commonUnitList, &unit)
		timestamp += 300
	}
	res := types.MeasureCommonUnitList{
		MeasureCommonUnitList: commonUnitList,
	}
	return &res
}
