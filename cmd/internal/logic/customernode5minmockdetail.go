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
	"strconv"
)

type CustomerNode5MinDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewCustomerNode5MinDetailLogic(ctx context.Context) *CustomerNode5MinDetailLogic {
	return &CustomerNode5MinDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *CustomerNode5MinDetailLogic) CustomerNode5MinMockDetail(mgo *config.Mongo, req *types.PostCustomerNode5MinReq) (resp *types.HadoopResp[types.MeasureCommonUnitList], err error) {
	resp = &types.HadoopResp[types.MeasureCommonUnitList]{} //初始化
	m := model.NewCustomerNode5MinModel(mgo.MongoUrl, mgo.MongoDatabase)
	filter := strconv.FormatInt(req.Start, 10) + "-" + strconv.FormatInt(req.End, 10)
	data, _ := m.FindByFilter(context.Background(), filter)
	if data == nil { //数据库中没有
		mockData := service.MockCustomerNode5Min(req.Start, req.End)
		fmt.Println(mockData)
		if mockData == nil {
			resp.Code = util.MockErr.Code
			resp.Error = util.MockErr.Msg
			return resp, errors.New("mock data nil")
		}
		//写入数据库
		CustomerNode5Min := types.CustomerNode5Min{
			Filter:                filter,
			MeasureCommonUnitList: *mockData,
		}
		err = m.Insert(context.Background(), &CustomerNode5Min)
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
		resp.Data = &data.MeasureCommonUnitList
		resp.Code = util.Success.Code
		resp.Error = ""
		return resp, nil
	}
}
