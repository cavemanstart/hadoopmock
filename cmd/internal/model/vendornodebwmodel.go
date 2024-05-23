package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	VendorNodeBwModel interface {
		Insert(ctx context.Context, data *common.NodeMomentDataList) error
		Update(ctx context.Context, data *common.NodeMomentDataList) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.NodeMomentDataList, error)
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

func (m *defaultVendorNodeBwModel) Insert(ctx context.Context, data *common.NodeMomentDataList) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultVendorNodeBwModel) Update(ctx context.Context, data *common.NodeMomentDataList) error {
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
func (m *defaultVendorNodeBwModel) FindById(ctx context.Context, id string) (*common.NodeMomentDataList, error) {
	filter := bson.M{
		"_id": id,
	}
	var res common.NodeMomentDataList
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
