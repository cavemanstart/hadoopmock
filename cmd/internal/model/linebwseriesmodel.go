package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	LineBWSeriesModel interface {
		Insert(ctx context.Context, data *types.LineBWSeries) error
		Update(ctx context.Context, data *types.LineBWSeries) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.LineBWSeries, error)
	}
	defaultLineBWSeriesModel struct {
		*mon.Model
	}
)

func NewLineBWSeriesModel(url string, db string) LineBWSeriesModel {
	return &defaultLineBWSeriesModel{
		Model: mon.MustNewModel(url, db, "vendorNodeMetric"),
	}
}

func (m *defaultLineBWSeriesModel) Insert(ctx context.Context, data *types.LineBWSeries) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultLineBWSeriesModel) Update(ctx context.Context, data *types.LineBWSeries) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultLineBWSeriesModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultLineBWSeriesModel) FindById(ctx context.Context, id string) (*types.LineBWSeries, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.LineBWSeries
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
