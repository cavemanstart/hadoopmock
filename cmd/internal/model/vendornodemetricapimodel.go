package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/hadoop"
)

type (
	VendorNodeMetricApiModel interface {
		Insert(ctx context.Context, data *hadoop.MeasureCommonData) error
		Update(ctx context.Context, data *hadoop.MeasureCommonData) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*hadoop.MeasureCommonData, error)
	}
	VendorNodeMetricModel struct {
		*mon.Model
	}
)

func NewVendorNodeMetricApiModel(url string, db string) VendorNodeMetricApiModel {
	return &VendorNodeMetricModel{
		Model: mon.MustNewModel(url, db, "vendornodemetric"),
	}
}

func (m *VendorNodeMetricModel) Insert(ctx context.Context, data *hadoop.MeasureCommonData) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *VendorNodeMetricModel) Update(ctx context.Context, data *hadoop.MeasureCommonData) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *VendorNodeMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *VendorNodeMetricModel) FindById(ctx context.Context, id string) (*hadoop.MeasureCommonData, error) {
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
