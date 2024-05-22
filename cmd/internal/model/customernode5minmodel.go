package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/hadoop"
)

type (
	CustomerNode5MinModel interface {
		Insert(ctx context.Context, data *hadoop.MeasureCommonUnitList) error
		Update(ctx context.Context, data *hadoop.MeasureCommonUnitList) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*hadoop.MeasureCommonUnitList, error)
	}
	defaultCustomerNode5MinModel struct {
		*mon.Model
	}
)

func NewCustomerNode5MinModel(url string, db string) CustomerNode5MinModel {
	return &defaultCustomerNode5MinModel{
		Model: mon.MustNewModel(url, db, "CustomerNode5min"),
	}
}

func (m *defaultCustomerNode5MinModel) Insert(ctx context.Context, data *hadoop.MeasureCommonUnitList) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultCustomerNode5MinModel) Update(ctx context.Context, data *hadoop.MeasureCommonUnitList) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultCustomerNode5MinModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultCustomerNode5MinModel) FindById(ctx context.Context, id string) (*hadoop.MeasureCommonUnitList, error) {
	filter := bson.M{
		"_id": id,
	}
	var res hadoop.MeasureCommonUnitList
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
