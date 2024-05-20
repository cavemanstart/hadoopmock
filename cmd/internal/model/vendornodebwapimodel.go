package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/hadoop"
)

type (
	VendorNodeBwApiModel interface {
		Insert(ctx context.Context, data *hadoop.VendorNodeMomentData) error
		Update(ctx context.Context, data *hadoop.VendorNodeMomentData) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*hadoop.VendorNodeMomentData, error)
	}
	VendorNodeBwModel struct {
		*mon.Model
	}
)

func NewVendorNodeBwApiModel(url string, db string) VendorNodeBwApiModel {
	return &VendorNodeBwModel{
		Model: mon.MustNewModel(url, db, "vendornodebw"),
	}
}

func (m *VendorNodeBwModel) Insert(ctx context.Context, data *hadoop.VendorNodeMomentData) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *VendorNodeBwModel) Update(ctx context.Context, data *hadoop.VendorNodeMomentData) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *VendorNodeBwModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *VendorNodeBwModel) FindById(ctx context.Context, id string) (*hadoop.VendorNodeMomentData, error) {
	filter := bson.M{
		"_id": id,
	}
	var res hadoop.VendorNodeMomentData
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
