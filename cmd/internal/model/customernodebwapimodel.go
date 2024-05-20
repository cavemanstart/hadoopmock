package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/hadoop"
)

type (
	CustomerNodeBwApiModel interface {
		Insert(ctx context.Context, data *hadoop.NodeMomentDataList) error
		Update(ctx context.Context, data *hadoop.NodeMomentDataList) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*hadoop.NodeMomentDataList, error)
	}
	CustomerNodeBwModel struct {
		*mon.Model
	}
)

func NewCustomerNodeBwApiModel(url string, db string) CustomerNodeBwApiModel {
	return &CustomerNodeBwModel{
		Model: mon.MustNewModel(url, db, "customernodebw"),
	}
}

func (m *CustomerNodeBwModel) Insert(ctx context.Context, data *hadoop.NodeMomentDataList) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *CustomerNodeBwModel) Update(ctx context.Context, data *hadoop.NodeMomentDataList) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *CustomerNodeBwModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *CustomerNodeBwModel) FindById(ctx context.Context, id string) (*hadoop.NodeMomentDataList, error) {
	filter := bson.M{
		"_id": id,
	}
	var res hadoop.NodeMomentDataList
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
