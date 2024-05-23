package main

import (
	"context"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/model"
)

func main() {
	mongoConfig := config.ReadMongoConfig()
	bwModel := model.NewVendorNodeBwModel(mongoConfig.MongoUrl, mongoConfig.MongoDatabase)
	_, err := bwModel.FindById(context.Background(), "test")
	if err != nil {
		return
	}
}
