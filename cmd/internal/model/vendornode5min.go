package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/hadoop"
)

type (
	VendorNode5MinApiModel interface {
		Insert(ctx context.Context, data *hadoop.MeasureCommonUnitList) error
		Update(ctx context.Context, data *hadoop.MeasureCommonUnitList) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*hadoop.MeasureCommonUnitList, error)
	}
	VendorNode5MinModel struct {
		*mon.Model
	}
)

func NewVendorNode5MinApiModel(url string, db string) VendorNode5MinApiModel {
	return &VendorNode5MinModel{
		Model: mon.MustNewModel(url, db, "vendornode5min"),
	}
}

func (m *VendorNode5MinModel) Insert(ctx context.Context, data *hadoop.MeasureCommonUnitList) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *VendorNode5MinModel) Update(ctx context.Context, data *hadoop.MeasureCommonUnitList) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *VendorNode5MinModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *VendorNode5MinModel) FindById(ctx context.Context, id string) (*hadoop.MeasureCommonUnitList, error) {
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
