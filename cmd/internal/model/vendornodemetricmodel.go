package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	VendorNodeMetricModel interface {
		Insert(ctx context.Context, data *types.VendorNodeMetric) error
		Update(ctx context.Context, data *types.VendorNodeMetric) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.VendorNodeMetric, error)
		FindByFilter(ctx context.Context, filter string) (*types.VendorNodeMetric, error)
	}
	defaultVendorNodeMetricModel struct {
		*mon.Model
	}
)

func NewVendorNodeMetricModel(url string, db string) VendorNodeMetricModel {
	return &defaultVendorNodeMetricModel{
		Model: mon.MustNewModel(url, db, "vendorNodeMetric"),
	}
}

func (m *defaultVendorNodeMetricModel) Insert(ctx context.Context, data *types.VendorNodeMetric) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultVendorNodeMetricModel) Update(ctx context.Context, data *types.VendorNodeMetric) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultVendorNodeMetricModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultVendorNodeMetricModel) FindById(ctx context.Context, id string) (*types.VendorNodeMetric, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.VendorNodeMetric
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
func (m *defaultVendorNodeMetricModel) FindByFilter(ctx context.Context, filter string) (*types.VendorNodeMetric, error) {
	var res types.VendorNodeMetric
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
