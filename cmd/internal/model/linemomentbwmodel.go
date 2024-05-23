package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/common"
)

type (
	LineMomentBandWidthModel interface {
		Insert(ctx context.Context, data *common.LineMomentBandWidth) error
		Update(ctx context.Context, data *common.LineMomentBandWidth) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*common.LineMomentBandWidth, error)
	}
	defaultLineMomentBandWidthModel struct {
		*mon.Model
	}
)

func NewLineMomentBandWidthModel(url string, db string) LineMomentBandWidthModel {
	return &defaultLineMomentBandWidthModel{
		Model: mon.MustNewModel(url, db, "vendorNodeMetric"),
	}
}

func (m *defaultLineMomentBandWidthModel) Insert(ctx context.Context, data *common.LineMomentBandWidth) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultLineMomentBandWidthModel) Update(ctx context.Context, data *common.LineMomentBandWidth) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultLineMomentBandWidthModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultLineMomentBandWidthModel) FindById(ctx context.Context, id string) (*common.LineMomentBandWidth, error) {
	filter := bson.M{
		"_id": id,
	}
	var res common.LineMomentBandWidth
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
