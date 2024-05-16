package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"hadoopmock/cmd/internal/pointerutil"

	"hadoopmock/cmd/internal/hadoop"
)

type (
	MeasureApiModel interface {
		Insert(ctx context.Context, data *hadoop.MeasureCommonData) error
		Update(ctx context.Context, data *hadoop.MeasureCommonData) error
		DeleteById(ctx context.Context, id string) error
		FindById(ctx context.Context, id string) (*hadoop.MeasureCommonData, error)
	}
	measureModel struct {
		*mon.Model
	}
)

func NewMeasureApiModel(url string, db string) MeasureApiModel {
	return &measureModel{
		Model: mon.MustNewModel(url, db, "measure_api"),
	}
}

func (m *measureModel) Insert(ctx context.Context, data *hadoop.MeasureCommonData) error {
	_, err := m.InsertOne(ctx, data)
	return err
}
func (m *measureModel) Update(ctx context.Context, data *hadoop.MeasureCommonData) error {
	filter := bson.M{
		"_id": data.Id,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data}, &mopt.UpdateOptions{Upsert: pointerutil.Pointer(true)})
	return err
}
func (m *measureModel) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}
	_, err := m.DeleteMany(ctx, filter)
	return err
}
func (m *measureModel) FindById(ctx context.Context, id string) (*hadoop.MeasureCommonData, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil
	}
	filter := bson.M{
		"_id": oid,
	}
	var res *hadoop.MeasureCommonData
	err = m.FindOne(ctx, res, filter)
	return res, err
}
