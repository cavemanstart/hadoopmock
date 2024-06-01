package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	VendorNodeBwModel interface {
		Insert(ctx context.Context, data *types.VendorNodeBw) error
		Update(ctx context.Context, data *types.VendorNodeBw) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.VendorNodeBw, error)
		FindByFilter(ctx context.Context, filter string) (*types.VendorNodeBw, error)
	}
	defaultVendorNodeBwModel struct {
		*mon.Model
	}
)

func NewVendorNodeBwModel(url string, db string) VendorNodeBwModel {
	return &defaultVendorNodeBwModel{
		Model: mon.MustNewModel(url, db, "vendorNodeBw"),
	}
}

func (m *defaultVendorNodeBwModel) Insert(ctx context.Context, data *types.VendorNodeBw) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultVendorNodeBwModel) Update(ctx context.Context, data *types.VendorNodeBw) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultVendorNodeBwModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultVendorNodeBwModel) FindById(ctx context.Context, id string) (*types.VendorNodeBw, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.VendorNodeBw
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
func (m *defaultVendorNodeBwModel) FindByFilter(ctx context.Context, filter string) (*types.VendorNodeBw, error) {
	var res types.VendorNodeBw
	err := m.FindOne(ctx, &res, bson.M{
		"filter": filter,
	})
	switch {
	case err == nil:
		return &res, nil
	case errors.Is(err, mon.ErrNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
