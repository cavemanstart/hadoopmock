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

type PingLossNodeRateDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewPingLossNodeRateDetailLogic(ctx context.Context) *PingLossNodeRateDetailLogic {
	return &PingLossNodeRateDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *PingLossNodeRateDetailLogic) PingLossNodeRateMockDetail(mgo *config.Mongo, req *types.PostPingLossNodeRateReq) (resp *types.HadoopResp[types.PingLossNodeRatioData], err error) {
	resp = &types.HadoopResp[types.PingLossNodeRatioData]{} //初始化
	m := model.NewPingLossNodeRatioModel(mgo.MongoUrl, mgo.MongoDatabase)
	filter := ""
	for _, nodeId := range req.NodeIds {
		filter += nodeId
	}
	filter += req.Start + req.End
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockPingLossNodeRate(req.NodeIds, req.Start, req.End)
		fmt.Println(mockData)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		PingLossNodeRate := types.PingLossNodeRatio{
			Filter:                filter,
			PingLossNodeRatioData: *mockData,
		}
		err = m.Insert(context.Background(), &PingLossNodeRate)
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
		resp.Data = &data.PingLossNodeRatioData
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	}
}
