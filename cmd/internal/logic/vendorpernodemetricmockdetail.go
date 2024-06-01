package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/model"
	"hadoopmock/cmd/internal/service"
	"hadoopmock/cmd/internal/types"
	"hadoopmock/cmd/internal/util"
	"strconv"
)

type VendorPerNodeMetricDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewVendorPerNodeMetricDetailLogic(ctx context.Context) *VendorPerNodeMetricDetailLogic {
	return &VendorPerNodeMetricDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *VendorPerNodeMetricDetailLogic) VendorPerNodeMetricMockDetail(mgo *config.Mongo, req *types.PostVendorPerNodeMetricReq) (resp *types.HadoopResp[types.MeasureCommonDataNodes], err error) {
	resp = &types.HadoopResp[types.MeasureCommonDataNodes]{} //初始化
	m := model.NewVendorPerNodeMetricModel(mgo.MongoUrl, mgo.MongoDatabase)
	nodeIds := []string{}
	filter := ""
	filter += req.Start + req.End + strconv.FormatInt(req.BillDays, 10)
	for _, node := range req.Nodes {
		nodeIds = append(nodeIds, node.NodeId)
		filter += node.NodeId
	}
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockVendorPerNodeMetric(req.Start, req.End, req.BillDays, nodeIds)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		vendorPerNodeMetric := types.VendorPerNodeMetric{
			Filter:                 filter,
			MeasureCommonDataNodes: *mockData,
		}
		err = m.Insert(context.Background(), &vendorPerNodeMetric)
		if err != nil {
			resp.Code = util.DbErr.Code
			resp.Error = util.DbErr.Msg
			return resp, err
		}
		resp.Data = mockData
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	} else { //数据库中已经存在
		resp.Data = &data.MeasureCommonDataNodes
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	}
}
