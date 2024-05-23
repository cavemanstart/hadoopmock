package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	LineBandWidthSeriesModel interface {
		Insert(ctx context.Context, data *common.LineBandWidthSeries) error
		Update(ctx context.Context, data *common.LineBandWidthSeries) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.LineBandWidthSeries, error)
	}
	defaultLineBandWidthSeriesModel struct {
		*mon.Model
	}
)

func NewLineBandWidthSeriesModel(url string, db string) LineBandWidthSeriesModel {
	return &defaultLineBandWidthSeriesModel{
		Model: mon.MustNewModel(url, db, "vendorNodeMetric"),
	}
}

func (m *defaultLineBandWidthSeriesModel) Insert(ctx context.Context, data *common.LineBandWidthSeries) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultLineBandWidthSeriesModel) Update(ctx context.Context, data *common.LineBandWidthSeries) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultLineBandWidthSeriesModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultLineBandWidthSeriesModel) FindById(ctx context.Context, id string) (*common.LineBandWidthSeries, error) {
	filter := bson.M{
		"_id": id,
	}
	var res common.LineBandWidthSeries
	err := m.FindOne(ctx, &res, filter)
	switch {
	case err == nil:
		return &res, nil
	case errors.Is(err, mon.ErrNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
