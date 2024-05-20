package dao

import (
	"context"
	"log"

	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
)

var (
	vendorNodeMetricApiModel model.VendorNodeMetricApiModel
)

func init() {
	mongoConfig := config.ReadConfig()
	vendorNodeMetricApiModel = model.NewVendorNodeMetricApiModel(mongoConfig.Mongo.Url, mongoConfig.Mongo.Database)
}
func InsertVendorNodeMetricModel(data *hadoop.MeasureCommonData) {
	resData := FindVendorNodeMetricModelApiById(data.Id)
	if resData == nil {
		err := vendorNodeMetricApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "measure api insert failed")
		}
		log.Println("measure api insert succeed")
	}
}
func DeleteVendorNodeMetricModel(id string) {
	err := vendorNodeMetricApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "measure api delete failed")
	}
	log.Println("measure api delete succeed")
}

func UpdateVendorNodeMetricModel(data *hadoop.MeasureCommonData) {
	err := vendorNodeMetricApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "measure api update failed")
	}
	log.Println("measure api update succeed")
}

func FindVendorNodeMetricModelApiById(id string) *hadoop.MeasureCommonData {
	resData, err := vendorNodeMetricApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, " measure api query failed")
	}
	log.Println(" measure api query succeed")
	return resData
}
