package dao

import (
	"context"
	"log"

	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
)

var (
	CustomerNode5MinApiModel model.CustomerNode5MinModel
)

func init() {
	mongoConfig := config.ReadConfig()
	CustomerNode5MinApiModel = model.NewCustomerNode5MinModel(mongoConfig.Mongo.Url, mongoConfig.Mongo.Database)
}
func InsertCustomerNode5MinModel(data *hadoop.MeasureCommonUnitList) {
	resData := FindCustomerNode5MinModelApiById(data.Id)
	if resData == nil {
		err := CustomerNode5MinApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "insert failed")
		}
		log.Println("insert succeed")
	}
}
func DeleteCustomerNode5MinModel(id string) {
	err := CustomerNode5MinApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "delete failed")
	}
	log.Println("delete succeed")
}

func UpdateCustomerNode5MinModel(data *hadoop.MeasureCommonUnitList) {
	err := CustomerNode5MinApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "update failed")
	}
	log.Println("update succeed")
}

func FindCustomerNode5MinModelApiById(id string) *hadoop.MeasureCommonUnitList {
	resData, err := CustomerNode5MinApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "query failed")
	}
	log.Println("query succeed")
	return resData
}
