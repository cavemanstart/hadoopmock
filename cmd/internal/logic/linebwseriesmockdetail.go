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

type LineBwSeriesDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewLineBwSeriesDetailLogic(ctx context.Context) *LineBwSeriesDetailLogic {
	return &LineBwSeriesDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *LineBwSeriesDetailLogic) LineBwSeriesDetailMockLogic(mgo *config.Mongo, req *types.PostLineBwSeriesReq) (resp *types.HadoopResp[types.LineBandWidthSeries], err error) {
	resp = &types.HadoopResp[types.LineBandWidthSeries]{} //初始化
	m := model.NewLineBwSeriesModel(mgo.MongoUrl, mgo.MongoDatabase)
	lines := []string{"line1", "line2"}
	filter := strconv.FormatInt(req.Start, 10) + strconv.FormatInt(req.End, 10)
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockLineBwSeries(lines, req.Start, req.End)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		LineBwSeries := types.LineBwSeries{
			Filter:              filter,
			LineBandWidthSeries: *mockData,
		}
		err = m.Insert(context.Background(), &LineBwSeries)
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
		resp.Data = &data.LineBandWidthSeries
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	}
}
