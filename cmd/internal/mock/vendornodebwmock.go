package mock

import (
	"hadoopmock/cmd/internal/common"
	"math/rand"
	"strconv"
	"time"
)

func MockVendorNodeBw(nodeIdList []string) *common.NodeMomentDataList {
	nodeMomentDatas := []*common.NodeMomentData{}
	for _, nodeId := range nodeIdList {
		tmpNodeMomentData := common.NodeMomentData{}
		tmpNodeMomentData.NodeId = nodeId
		tmpNodeMomentData.Bandwidth = baseBandwidth + rand.Int63n(rangeBandwidth)
		nodeMomentDatas = append(nodeMomentDatas, &tmpNodeMomentData)
	}
	nodeMomentDataList := common.NodeMomentDataList{
		Id:                 idPrefix + strconv.FormatInt(time.Now().UnixMilli()+rand.Int63n(10000), 10),
		NodeMomentDataList: nodeMomentDatas,
	}
	return &nodeMomentDataList
}
