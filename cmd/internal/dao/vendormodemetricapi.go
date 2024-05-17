package dao

import (
	"context"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
	"log"
)

var (
	mongoUrl                        = "mongodb://localhost:27017"
	mongoDb                  string = "hadoop"
	vendorNodeMetricApiModel model.VendorNodeMetricApiModel
)

func InsertVendorNodeMetricModel(data *hadoop.MeasureCommonData) {
	if vendorNodeMetricApiModel == nil {
		vendorNodeMetricApiModel = model.NewVendorNodeMetricApiModel(mongoUrl, mongoDb)
	}
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
	if vendorNodeMetricApiModel == nil {
		vendorNodeMetricApiModel = model.NewVendorNodeMetricApiModel(mongoUrl, mongoDb)
	}
	err := vendorNodeMetricApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "measure api delete failed")
	}
	log.Println("measure api delete succeed")
}

func UpdateVendorNodeMetricModel(data *hadoop.MeasureCommonData) {
	if vendorNodeMetricApiModel == nil {
		vendorNodeMetricApiModel = model.NewVendorNodeMetricApiModel(mongoUrl, mongoDb)
	}
	err := vendorNodeMetricApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "measure api update failed")
	}
	log.Println("measure api update succeed")
}

func FindVendorNodeMetricModelApiById(id string) *hadoop.MeasureCommonData {
	if vendorNodeMetricApiModel == nil {
		vendorNodeMetricApiModel = model.NewVendorNodeMetricApiModel(mongoUrl, mongoDb)
	}
	resData, err := vendorNodeMetricApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, " measure api query failed")
	}
	log.Println(" measure api query succeed")
	return resData
}
