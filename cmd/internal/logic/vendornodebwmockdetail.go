package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/model"
	"hadoopmock/cmd/internal/service"
	"hadoopmock/cmd/internal/types"
	"hadoopmock/cmd/internal/util"
)

type VendorNodeBwDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewVendorNodeBwDetailLogic(ctx context.Context) *VendorNodeBwDetailLogic {
	return &VendorNodeBwDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *VendorNodeBwDetailLogic) VendorNodeBwMockDetail(mgo *config.Mongo, req *types.PostVendorNodeBwReq) (resp *types.HadoopResp[types.NodeMomentDataList], err error) {
	resp = &types.HadoopResp[types.NodeMomentDataList]{} //初始化
	m := model.NewVendorNodeBwModel(mgo.MongoUrl, mgo.MongoDatabase)
	nodeIds := []string{}
	filter := ""
	for _, node := range req.Nodes {
		nodeIds = append(nodeIds, node.NodeId)
		filter += node.NodeId
	}
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockVendorNodeBw(nodeIds)
		fmt.Println(mockData)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		vendorNodeBw := types.VendorNodeBw{
			Filter:             filter,
			NodeMomentDataList: *mockData,
		}
		err = m.Insert(context.Background(), &vendorNodeBw)
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
		resp.Data = &data.NodeMomentDataList
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	}
}
