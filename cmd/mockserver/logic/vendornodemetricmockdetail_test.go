package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"hadoopmock/cmd/mockserver/config"
	"hadoopmock/cmd/mockserver/types"
)

var (
	sc *config.ServiceContext = config.NewServiceContext(config.ReadConfig("../config.yaml"))
)

func TestVendorNodeMetricDetailLogic_VendorNodeMetricMockDetail(t *testing.T) {
	node1 := types.MeasureMetricNode{NodeId: "test001"}
	nodes := []*types.MeasureMetricNode{&node1}
	req := &types.PostVendorNodeMetricReq{
		Nodes:       nodes,
		Start:       "2023-11-01",
		End:         "2023-11-08",
		BillDays:    int64(6),
		MeasureType: "peak95",
	}
	l := NewVendorNodeMetricDetailLogic(sc.Context)
	resp, err := l.VendorNodeMetricMockDetail(sc, req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
