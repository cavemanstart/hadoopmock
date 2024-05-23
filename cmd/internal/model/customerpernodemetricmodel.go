package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	CustomerPerNodeMetricModel interface {
		Insert(ctx context.Context, data *common.MeasureCommonDataPerNode) error
		Update(ctx context.Context, data *common.MeasureCommonDataPerNode) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.MeasureCommonDataPerNode, error)
	}
	defaultCustomerPerNodeMetricModel struct {
		*mon.Model
	}
)

func NewCustomerPerNodeMetricModel(url string, db string) CustomerPerNodeMetricModel {
	return &defaultCustomerPerNodeMetricModel{
		Model: mon.MustNewModel(url, db, "CustomerPerNodeMetric"),
	}
}

func (m *defaultCustomerPerNodeMetricModel) Insert(ctx context.Context, data *common.MeasureCommonDataPerNode) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultCustomerPerNodeMetricModel) Update(ctx context.Context, data *common.MeasureCommonDataPerNode) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultCustomerPerNodeMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultCustomerPerNodeMetricModel) FindById(ctx context.Context, id string) (*common.MeasureCommonDataPerNode, error) {
	filter := bson.M{
		"_id": id,
	}
	var res common.MeasureCommonDataPerNode
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