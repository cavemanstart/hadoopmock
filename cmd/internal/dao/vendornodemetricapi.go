package dao

import (
	"context"
	"log"

	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
)

var (
	vendorNodeMetricApiModel model.VendorNodeMetricModel
)

func init() {
	mongoConfig := config.ReadConfig()
	vendorNodeMetricApiModel = model.NewVendorNodeMetricModel(mongoConfig.Mongo.Url, mongoConfig.Mongo.Database)
}
func InsertVendorNodeMetricModel(data *hadoop.MeasureCommonData) {
	resData := FindVendorNodeMetricModelApiById(data.Id)
	if resData == nil {
		err := vendorNodeMetricApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "insert failed")
		}
		log.Println("insert succeed")
	}
}
func DeleteVendorNodeMetricModel(id string) {
	err := vendorNodeMetricApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "delete failed")
	}
	log.Println("delete succeed")
}

func UpdateVendorNodeMetricModel(data *hadoop.MeasureCommonData) {
	err := vendorNodeMetricApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "update failed")
	}
	log.Println("update succeed")
}

func FindVendorNodeMetricModelApiById(id string) *hadoop.MeasureCommonData {
	resData, err := vendorNodeMetricApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "query failed")
	}
	log.Println("query succeed")
	return resData
}
