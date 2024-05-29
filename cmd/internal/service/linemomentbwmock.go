package service

import (
	"hadoopmock/cmd/internal/types"
	"math/rand"
)

func MockLineMomentBw(nodeLinesMap map[string][]string) *types.LineMomentBandWidth {
	lineMoment := map[string][]*types.LineCommonUint{}
	for nodeId, lines := range nodeLinesMap {
		slice := []*types.LineCommonUint{}
		for _, line := range lines {
			unit := types.LineCommonUint{
				Line:      line,
				Bandwidth: baseBandwidth + rand.Int63n(rangeBandwidth),
			}
			slice = append(slice, &unit)
		}
		lineMoment[nodeId] = slice
	}
	return &types.LineMomentBandWidth{
		LineMoment: lineMoment,
	}
}
