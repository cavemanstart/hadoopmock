package logic

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hadoopmock/cmd/mockserver/types"
	"testing"
	"time"
)

func TestGenerateMeasureDetailLogic_GenerateMeasureDetail(t *testing.T) {
	nodeId, vendorId, customerId := "test001", "1813", "12321"
	start := time.Date(2023, 11, 01, 0, 0, 0, 0, time.Local)
	end := time.Date(2023, 11, 18, 0, 0, 0, 0, time.Local)
	var flow = int64(2024000000)
	req := &types.PostGenerateMeasureReq{
		NodeId:     nodeId,
		VendorId:   vendorId,
		CustomerId: customerId,
		Start:      start,
		End:        end,
		Flow:       flow,
	}
	l := NewGenerateMeasureDetailLogic(context.Background())
	err := l.GenerateMeasureDetail(sc, req)
	assert.Nil(t, err)
}
