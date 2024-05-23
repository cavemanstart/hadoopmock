package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	CustomerNodeBwModel interface {
		Insert(ctx context.Context, data *common.NodeMomentDataList) error
		Update(ctx context.Context, data *common.NodeMomentDataList) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.NodeMomentDataList, error)
	}
	defaultCustomerNodeBwModel struct {
		*mon.Model
	}
)

func NewCustomerNodeBwModel(url string, db string) CustomerNodeBwModel {
	return &defaultCustomerNodeBwModel{
		Model: mon.MustNewModel(url, db, "customerNodeBw"),
	}
}

func (m *defaultCustomerNodeBwModel) Insert(ctx context.Context, data *common.NodeMomentDataList) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultCustomerNodeBwModel) Update(ctx context.Context, data *common.NodeMomentDataList) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultCustomerNodeBwModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultCustomerNodeBwModel) FindById(ctx context.Context, id string) (*common.NodeMomentDataList, error) {
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
