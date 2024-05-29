package service

import (
	"hadoopmock/cmd/internal/types"
	"math/rand"
)

func MockCustomerNodeBW(nodeIdList []string) *types.NodeMomentDataList {
	nodeMomentDatas := []*types.NodeMomentData{}
	for _, nodeId := range nodeIdList {
		tmpNodeMomentData := types.NodeMomentData{}
		tmpNodeMomentData.NodeId = nodeId
		tmpNodeMomentData.Bandwidth = baseBandwidth + rand.Int63n(rangeBandwidth)
		nodeMomentDatas = append(nodeMomentDatas, &tmpNodeMomentData)
	}
	nodeMomentDataList := types.NodeMomentDataList{
		NodeMomentDataList: nodeMomentDatas,
	}
	return &nodeMomentDataList
}
