package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	VendorNode5MinModel interface {
		Insert(ctx context.Context, data *types.VendorNode5Min) error
		Update(ctx context.Context, data *types.VendorNode5Min) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.VendorNode5Min, error)
	}
	defaultVendorNode5MinModel struct {
		*mon.Model
	}
)

func NewVendorNode5MinModel(url string, db string) VendorNode5MinModel {
	return &defaultVendorNode5MinModel{
		Model: mon.MustNewModel(url, db, "vendorNode5min"),
	}
}

func (m *defaultVendorNode5MinModel) Insert(ctx context.Context, data *types.VendorNode5Min) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultVendorNode5MinModel) Update(ctx context.Context, data *types.VendorNode5Min) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultVendorNode5MinModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultVendorNode5MinModel) FindById(ctx context.Context, id string) (*types.VendorNode5Min, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.VendorNode5Min
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
