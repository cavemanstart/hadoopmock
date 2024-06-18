package config

import (
	"context"

	"hadoopmock/cmd/mockserver/model"
)

type ServiceContext struct {
	Context                   context.Context
	JarvisChargeNode5minModel model.JarvisChargeNode5minModel
}

func NewServiceContext(cfg *Config) *ServiceContext {
	return &ServiceContext{
		Context:                   context.Background(),
		JarvisChargeNode5minModel: model.NewJarvisChargeNode5minModel(cfg.MongoUrl, cfg.MongoDatabase),
	}
}
