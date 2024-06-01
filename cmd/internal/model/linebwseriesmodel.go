package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	LineBwSeriesModel interface {
		Insert(ctx context.Context, data *types.LineBwSeries) error
		Update(ctx context.Context, data *types.LineBwSeries) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.LineBwSeries, error)
		FindByFilter(ctx context.Context, filter string) (*types.LineBwSeries, error)
	}
	defaultLineBwSeriesModel struct {
		*mon.Model
	}
)

func NewLineBwSeriesModel(url string, db string) LineBwSeriesModel {
	return &defaultLineBwSeriesModel{
		Model: mon.MustNewModel(url, db, "lineBwSeries"),
	}
}

func (m *defaultLineBwSeriesModel) Insert(ctx context.Context, data *types.LineBwSeries) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultLineBwSeriesModel) Update(ctx context.Context, data *types.LineBwSeries) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultLineBwSeriesModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultLineBwSeriesModel) FindById(ctx context.Context, id string) (*types.LineBwSeries, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.LineBwSeries
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
func (m *defaultLineBwSeriesModel) FindByFilter(ctx context.Context, filter string) (*types.LineBwSeries, error) {
	var res types.LineBwSeries
	err := m.FindOne(ctx, &res, bson.M{
		"filter": filter,
	})
	switch {
	case err == nil:
		return &res, nil
	case errors.Is(err, mon.ErrNotFound):
		return nil, nil
	default:
		return nil, err
	}
}
