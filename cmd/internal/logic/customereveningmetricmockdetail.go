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

type CustomerEveningMetricDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewCustomerEveningMetricDetailLogic(ctx context.Context) *CustomerEveningMetricDetailLogic {
	return &CustomerEveningMetricDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *CustomerEveningMetricDetailLogic) CustomerEveningMetricMockDetail(mgo *config.Mongo, req *types.PostCustomerEveningMetricReq) (resp *types.HadoopResp[types.MeasureCommonData], err error) {
	resp = &types.HadoopResp[types.MeasureCommonData]{} //初始化
	m := model.NewCustomerEveningMetricModel(mgo.MongoUrl, mgo.MongoDatabase)
	filter := req.Start + "-" + req.End
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockCustomerEveningMetric(req.Start, req.End)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		CustomerEveningMetric := types.CustomerEveningMetric{
			Filter:            filter,
			MeasureCommonData: *mockData,
		}
		err = m.Insert(context.Background(), &CustomerEveningMetric)
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
		resp.Data = &data.MeasureCommonData
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	}
}
