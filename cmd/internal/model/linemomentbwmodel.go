package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"hadoopmock/cmd/internal/types"
)

type (
	LineMomentBwModel interface {
		Insert(ctx context.Context, data *types.LineMomentBw) error
		Update(ctx context.Context, data *types.LineMomentBw) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*types.LineMomentBw, error)
		FindByFilter(ctx context.Context, filter string) (*types.LineMomentBw, error)
	}
	defaultLineMomentBwModel struct {
		*mon.Model
	}
)

func NewLineMomentBwModel(url string, db string) LineMomentBwModel {
	return &defaultLineMomentBwModel{
		Model: mon.MustNewModel(url, db, "lineMomentBw"),
	}
}

func (m *defaultLineMomentBwModel) Insert(ctx context.Context, data *types.LineMomentBw) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *defaultLineMomentBwModel) Update(ctx context.Context, data *types.LineMomentBw) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}
func (m *defaultLineMomentBwModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *defaultLineMomentBwModel) FindById(ctx context.Context, id string) (*types.LineMomentBw, error) {
	filter := bson.M{
		"_id": id,
	}
	var res types.LineMomentBw
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
func (m *defaultLineMomentBwModel) FindByFilter(ctx context.Context, filter string) (*types.LineMomentBw, error) {
	var res types.LineMomentBw
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
