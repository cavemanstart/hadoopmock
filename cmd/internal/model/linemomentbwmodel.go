package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	LineMomentBWModel interface {
		Insert(ctx context.Context, data *types.LineMomentBW) error
		Update(ctx context.Context, data *types.LineMomentBW) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.LineMomentBW, error)
	}
	defaultLineMomentBWModel struct {
		*mon.Model
	}
)

func NewLineMomentBWModel(url string, db string) LineMomentBWModel {
	return &defaultLineMomentBWModel{
		Model: mon.MustNewModel(url, db, "vendorNodeMetric"),
	}
}

func (m *defaultLineMomentBWModel) Insert(ctx context.Context, data *types.LineMomentBW) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultLineMomentBWModel) Update(ctx context.Context, data *types.LineMomentBW) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultLineMomentBWModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultLineMomentBWModel) FindById(ctx context.Context, id string) (*types.LineMomentBW, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.LineMomentBW
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
