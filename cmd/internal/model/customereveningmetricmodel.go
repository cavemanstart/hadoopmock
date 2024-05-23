package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	CustomerEveningMetricModel interface {
		Insert(ctx context.Context, data *common.MeasureCommonData) error
		Update(ctx context.Context, data *common.MeasureCommonData) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.MeasureCommonData, error)
	}
	defaultCustomerEveningMetricModel struct {
		*mon.Model
	}
)

func NewCustomerEveningMetricModel(url string, db string) CustomerEveningMetricModel {
	return &defaultCustomerEveningMetricModel{
		Model: mon.MustNewModel(url, db, "CustomerEveningMetric"),
	}
}

func (m *defaultCustomerEveningMetricModel) Insert(ctx context.Context, data *common.MeasureCommonData) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultCustomerEveningMetricModel) Update(ctx context.Context, data *common.MeasureCommonData) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultCustomerEveningMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultCustomerEveningMetricModel) FindById(ctx context.Context, id string) (*common.MeasureCommonData, error) {
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
