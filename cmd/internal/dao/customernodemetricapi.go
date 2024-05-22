package dao

import (
	"context"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
	"log"
)

var (
	CustomerNodeMetricApiModel model.CustomerNodeMetricApiModel
)

func init() {
	mongoConfig := config.ReadConfig()
	CustomerNodeMetricApiModel = model.NewCustomerNodeMetricModel(mongoConfig.Mongo.Url, mongoConfig.Mongo.Database)
}
func InsertCustomerNodeMetricModel(data *hadoop.MeasureCommonData) {
	resData := FindCustomerNodeMetricModelApiById(data.Id)
	if resData == nil {
		err := CustomerNodeMetricApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "insert failed")
		}
		log.Println("insert succeed")
	}
}
func DeleteCustomerNodeMetricModel(id string) {
	err := CustomerNodeMetricApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "delete failed")
	}
	log.Println("delete succeed")
}

func UpdateCustomerNodeMetricModel(data *hadoop.MeasureCommonData) {
	err := CustomerNodeMetricApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "update failed")
	}
	log.Println("update succeed")
}

func FindCustomerNodeMetricModelApiById(id string) *hadoop.MeasureCommonData {
	resData, err := CustomerNodeMetricApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, " query failed")
	}
	log.Println(" query succeed")
	return resData
}
