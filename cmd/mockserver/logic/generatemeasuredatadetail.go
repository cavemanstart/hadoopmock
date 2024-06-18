package logic

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"hadoopmock/cmd/mockserver/config"
	"hadoopmock/cmd/mockserver/model"
	"hadoopmock/cmd/mockserver/types"
)

type GenerateMeasureDetailLogic struct {
	logx.Logger
	ctx context.Context
}

func NewGenerateMeasureDetailLogic(ctx context.Context) *GenerateMeasureDetailLogic {
	return &GenerateMeasureDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}
func (l *GenerateMeasureDetailLogic) GenerateMeasureDetail(sc *config.ServiceContext, req *types.PostGenerateMeasureReq) error {
	startDate := req.Start
	endDate := req.End
	if endDate.Before(startDate) {
		logx.Errorf("invalid args: %v > %v", startDate, endDate)
		return errors.New("invalid args")
	}
	err := MockBigDataBillingV2Data(sc, req)
	if err != nil {
		return err
	}
	return nil
}

func MockBigDataBillingV2Data(sc *config.ServiceContext, req *types.PostGenerateMeasureReq) error {
	logx.SetLevel(logx.ErrorLevel)
	defer logx.SetLevel(logx.DebugLevel)
	start, end, nodeId, vendorId, customerId, flow := req.Start, req.End, req.NodeId, req.VendorId, req.CustomerId, req.Flow
	for date := start; date.Before(end); date = date.AddDate(0, 0, 1) {
		var (
			times   []time.Time
			inFlow  []int64
			outFlow []int64
		)
		for date5min := date; date5min.Before(date.AddDate(0, 0, 1)); date5min = date5min.Add(time.Duration(5) * time.Minute) {
			times = append(times, date5min)
			inFlow = append(inFlow, 18750000000)
			outFlow = append(outFlow, flow)
		}
		jarvisChargeNode5min := &model.JarvisChargeNode5min{
			NodeId:  nodeId,
			Uid:     customerId,
			Vendor:  vendorId,
			Day:     date,
			Times:   times,
			InFlow:  inFlow,
			OutFlow: outFlow,
		}
		err := sc.JarvisChargeNode5minModel.Insert(sc.Context, jarvisChargeNode5min)
		if err != nil {
			return err
		}
	}
	return nil
}
