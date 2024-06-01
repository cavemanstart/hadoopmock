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
)

type CustomerNodeBwDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewCustomerNodeBwDetailLogic(ctx context.Context) *CustomerNodeBwDetailLogic {
	return &CustomerNodeBwDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *CustomerNodeBwDetailLogic) CustomerNodeBwMockDetail(mgo *config.Mongo, req *types.PostCustomerNodeBwReq) (resp *types.HadoopResp[types.NodeMomentDataList], err error) {
	resp = &types.HadoopResp[types.NodeMomentDataList]{} //初始化
	m := model.NewCustomerNodeBwModel(mgo.MongoUrl, mgo.MongoDatabase)
	nodeIds := []string{}
	filter := ""
	for _, node := range req.Nodes {
		nodeIds = append(nodeIds, node.NodeId)
		filter += node.NodeId
	}
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockCustomerNodeBW(nodeIds)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		CustomerNodeBw := types.CustomerNodeBw{
			Filter:             filter,
			NodeMomentDataList: *mockData,
		}
		err = m.Insert(context.Background(), &CustomerNodeBw)
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
