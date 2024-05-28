package mock

import (
	"hadoopmock/cmd/internal/common"
	"math/rand"
	"strconv"
	"time"
)

func MockVendorNode5Min(startTime int64, endTime int64) *common.MeasureCommonUnitList {
	commonUnitList := []*common.MeasureCommonUnit{}
	timestamp := startTime
	for {
		if timestamp >= endTime {
			break
		}
		unit := common.MeasureCommonUnit{
			Bandwidth: baseBandwidth + rand.Int63n(rangeBandwidth),
			Time:      timestamp,
		}
		commonUnitList = append(commonUnitList, &unit)
		timestamp += 300
	}
	res := common.MeasureCommonUnitList{
		Id:                    idPrefix + strconv.FormatInt(time.Now().UnixMilli()+rand.Int63n(10000), 10),
		MeasureCommonUnitList: commonUnitList,
	}
	return &res
}
