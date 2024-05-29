package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	VendorPerNodeMetricModel interface {
		Insert(ctx context.Context, data *types.VendorPerNodeMetric) error
		Update(ctx context.Context, data *types.VendorPerNodeMetric) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.VendorPerNodeMetric, error)
	}
	defaultVendorPerNodeMetricModel struct {
		*mon.Model
	}
)

func NewVendorPerNodeMetricModel(url string, db string) VendorPerNodeMetricModel {
	return &defaultVendorPerNodeMetricModel{
		Model: mon.MustNewModel(url, db, "vendorPerNodeMetric"),
	}
}

func (m *defaultVendorPerNodeMetricModel) Insert(ctx context.Context, data *types.VendorPerNodeMetric) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultVendorPerNodeMetricModel) Update(ctx context.Context, data *types.VendorPerNodeMetric) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultVendorPerNodeMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultVendorPerNodeMetricModel) FindById(ctx context.Context, id string) (*types.VendorPerNodeMetric, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.VendorPerNodeMetric
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
