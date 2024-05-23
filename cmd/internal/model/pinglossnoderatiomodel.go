package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	PingLossNodeRatioModel interface {
		Insert(ctx context.Context, data *common.PingLossNodeRatio) error
		Update(ctx context.Context, data *common.PingLossNodeRatio) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.PingLossNodeRatio, error)
	}
	defaultPingLossNodeRatioModel struct {
		*mon.Model
	}
)

func NewPingLossNodeRatioModel(url string, db string) PingLossNodeRatioModel {
	return &defaultPingLossNodeRatioModel{
		Model: mon.MustNewModel(url, db, "vendorNodeMetric"),
	}
}

func (m *defaultPingLossNodeRatioModel) Insert(ctx context.Context, data *common.PingLossNodeRatio) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultPingLossNodeRatioModel) Update(ctx context.Context, data *common.PingLossNodeRatio) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultPingLossNodeRatioModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultPingLossNodeRatioModel) FindById(ctx context.Context, id string) (*common.PingLossNodeRatio, error) {
	filter := bson.M{
		"_id": id,
	}
	var res common.PingLossNodeRatio
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
