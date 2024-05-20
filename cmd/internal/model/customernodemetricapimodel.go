package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/hadoop"
)

type (
	CustomerNodeMetricApiModel interface {
		Insert(ctx context.Context, data *hadoop.MeasureCommonData) error
		Update(ctx context.Context, data *hadoop.MeasureCommonData) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*hadoop.MeasureCommonData, error)
	}
	CustomerNodeMetricModel struct {
		*mon.Model
	}
)

func NewCustomerNodeMetricApiModel(url string, db string) CustomerNodeMetricApiModel {
	return &CustomerNodeMetricModel{
		Model: mon.MustNewModel(url, db, "customernodemetric"),
	}
}

func (m *CustomerNodeMetricModel) Insert(ctx context.Context, data *hadoop.MeasureCommonData) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *CustomerNodeMetricModel) Update(ctx context.Context, data *hadoop.MeasureCommonData) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *CustomerNodeMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *CustomerNodeMetricModel) FindById(ctx context.Context, id string) (*hadoop.MeasureCommonData, error) {
	filter := bson.M{
		"_id": id,
	}
	var res hadoop.MeasureCommonData
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
