package dao

import (
	"context"
	"log"

	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
)

var (
	CustomerNodeBwApiModel model.CustomerNodeBwApiModel
)

func init() {
	mongoConfig := config.ReadConfig()
	CustomerNodeBwApiModel = model.NewCustomerNodeBwApiModel(mongoConfig.Mongo.Url, mongoConfig.Mongo.Database)
}
func InsertCustomerNodeBwModel(data *hadoop.NodeMomentDataList) {
	resData := FindCustomerNodeBwModelApiById(data.Id)
	if resData == nil {
		err := CustomerNodeBwApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "insert failed")
		}
		log.Println("insert succeed")
	}
}
func DeleteCustomerNodeBwModel(id string) {
	err := CustomerNodeBwApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "delete failed")
	}
	log.Println("delete succeed")
}

func UpdateCustomerNodeBwModel(data *hadoop.NodeMomentDataList) {
	err := CustomerNodeBwApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "update failed")
	}
	log.Println("update succeed")
}

func FindCustomerNodeBwModelApiById(id string) *hadoop.NodeMomentDataList {
	resData, err := CustomerNodeBwApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, " query failed")
	}
	log.Println(" query succeed")
	return resData
}
