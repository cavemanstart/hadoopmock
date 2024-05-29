package service

import (
	"hadoopmock/cmd/internal/types"
	"math/rand"
)

func MockLineBwSeries(lines []string, startTime int64, endTime int64) *types.LineBandWidthSeries {
	lineBandWidth := map[string][]*types.MeasureCommonUnit{}
	for _, line := range lines {
		slice := []*types.MeasureCommonUnit{}
		for t := startTime; t < endTime; t += 300 {
			unit := types.MeasureCommonUnit{
				Bandwidth: baseBandwidth + rand.Int63n(rangeBandwidth),
				Time:      t,
			}
			slice = append(slice, &unit)
		}
		lineBandWidth[line] = slice
	}
	return &types.LineBandWidthSeries{
		LineBandWidth: lineBandWidth,
	}
}
