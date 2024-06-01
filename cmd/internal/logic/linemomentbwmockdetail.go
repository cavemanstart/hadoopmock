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

type LineMomentBwDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewLineMomentBwDetailLogic(ctx context.Context) *LineMomentBwDetailLogic {
	return &LineMomentBwDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *LineMomentBwDetailLogic) LineMomentBwDetailMockLogic(mgo *config.Mongo, req *types.PostLineMomentBwReq) (resp *types.HadoopResp[types.LineMomentBandWidth], err error) {
	resp = &types.HadoopResp[types.LineMomentBandWidth]{} //初始化
	m := model.NewLineMomentBwModel(mgo.MongoUrl, mgo.MongoDatabase)
	nodeLinesMap := map[string][]string{}
	filter := ""
	for _, nodeId := range req.NodeIds {
		nodeLinesMap[nodeId] = []string{
			"line1",
			"line2",
		}
		filter += nodeId
	}
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockLineMomentBw(nodeLinesMap)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		LineMomentBw := types.LineMomentBw{
			Filter:              filter,
			LineMomentBandWidth: *mockData,
		}
		err = m.Insert(context.Background(), &LineMomentBw)
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
		resp.Data = &data.LineMomentBandWidth
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	}
}
