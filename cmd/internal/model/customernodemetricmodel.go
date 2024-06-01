package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	CustomerNodeMetricModel interface {
		Insert(ctx context.Context, data *types.CustomerNodeMetric) error
		Update(ctx context.Context, data *types.CustomerNodeMetric) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.CustomerNodeMetric, error)
		FindByFilter(ctx context.Context, filter string) (*types.CustomerNodeMetric, error)
	}
	defaultCustomerNodeMetricModel struct {
		*mon.Model
	}
)

func NewCustomerNodeMetricModel(url string, db string) CustomerNodeMetricModel {
	return &defaultCustomerNodeMetricModel{
		Model: mon.MustNewModel(url, db, "customerNodeMetric"),
	}
}

func (m *defaultCustomerNodeMetricModel) Insert(ctx context.Context, data *types.CustomerNodeMetric) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultCustomerNodeMetricModel) Update(ctx context.Context, data *types.CustomerNodeMetric) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultCustomerNodeMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultCustomerNodeMetricModel) FindById(ctx context.Context, id string) (*types.CustomerNodeMetric, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.CustomerNodeMetric
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
func (m *defaultCustomerNodeMetricModel) FindByFilter(ctx context.Context, filter string) (*types.CustomerNodeMetric, error) {
	var res types.CustomerNodeMetric
	err := m.FindOne(ctx, &res, bson.M{"filter": filter})
	switch {
	case err == nil:
		return &res, nil
	case errors.Is(err, mon.ErrNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
