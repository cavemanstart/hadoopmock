package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	VendorEveningMetricModel interface {
		Insert(ctx context.Context, data *common.MeasureCommonData) error
		Update(ctx context.Context, data *common.MeasureCommonData) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.MeasureCommonData, error)
	}
	defaultVendorEveningMetricModel struct {
		*mon.Model
	}
)

func NewVendorEveningMetricModel(url string, db string) VendorEveningMetricModel {
	return &defaultVendorEveningMetricModel{
		Model: mon.MustNewModel(url, db, "vendorEveningMetric"),
	}
}

func (m *defaultVendorEveningMetricModel) Insert(ctx context.Context, data *common.MeasureCommonData) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultVendorEveningMetricModel) Update(ctx context.Context, data *common.MeasureCommonData) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultVendorEveningMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultVendorEveningMetricModel) FindById(ctx context.Context, id string) (*common.MeasureCommonData, error) {
	filter := bson.M{
		"_id": id,
	}
	var res common.MeasureCommonData
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
