package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	JarvisChargeNode5min struct {
		NodeId  string      `bson:"node_id,omitempty"   json:"node_id,omitempty"` //节点ID
		Uid     string      `bson:"uid,omitempty"       json:"uid,omitempty"`     // 业务方 Id
		Vendor  string      `bson:"vendor,omitempty"   json:"vendor,omitempty"`
		Day     time.Time   `bson:"day,omitempty"   json:"day,omitempty"`
		Times   []time.Time `bson:"times,omitempty"   json:"times,omitempty"`
		InFlow  []int64     `bson:"in_flow,omitempty"  json:"in_flow,omitempty"`
		OutFlow []int64     `bson:"out_flow,omitempty"  json:"out_flow,omitempty"`
	}

	JarvisChargeNode5minModel interface {
		Insert(ctx context.Context, data *JarvisChargeNode5min) error
		Update(ctx context.Context, data *JarvisChargeNode5min) error
		Delete(ctx context.Context, id string) error
		FindByNodeIdAndDate(ctx context.Context, nodeId string, day time.Time) (data *JarvisChargeNode5min, err error)
	}

	defaultJarvisChargeNode5minModel struct {
		*mon.Model
	}
)

func NewJarvisChargeNode5minModel(url string, db string) JarvisChargeNode5minModel {
	return &defaultJarvisChargeNode5minModel{
		Model: mon.MustNewModel(url, db, "jarvis_charge_node_5min"),
	}
}

func (m *defaultJarvisChargeNode5minModel) Insert(ctx context.Context, data *JarvisChargeNode5min) (err error) {
	_, err = m.InsertOne(ctx, data)
	return err
}

func (m *defaultJarvisChargeNode5minModel) Delete(ctx context.Context, id string) error {
	filter := bson.M{
		"node_id": bson.M{"$regex": id},
	}
	_, err := m.DeleteMany(ctx, filter)
	return err

}

func (m *defaultJarvisChargeNode5minModel) Update(ctx context.Context, data *JarvisChargeNode5min) error {
	filter := bson.M{
		"node_id": data.NodeId,
	}
	_, err := m.UpdateOne(ctx, filter, bson.M{"$set": data})
	return err
}

func (m *defaultJarvisChargeNode5minModel) FindByNodeIdAndDate(ctx context.Context, nodeId string, day time.Time) (*JarvisChargeNode5min, error) {
	data := &JarvisChargeNode5min{}
	err := m.FindOne(ctx, data, bson.M{
		"node_id": nodeId,
		"day":     day,
	})
	return data, err
}
