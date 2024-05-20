package dao

import (
	"context"
	"log"

	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
)

var (
	vendorNode5MinApiModel model.VendorNode5MinApiModel
)

func init() {
	mongoConfig := config.ReadConfig()
	vendorNode5MinApiModel = model.NewVendorNode5MinApiModel(mongoConfig.Mongo.Url, mongoConfig.Mongo.Database)
}
func InsertVendorNode5MinModel(data *hadoop.MeasureCommonUnitList) {
	resData := FindVendorNode5MinModelApiById(data.Id)
	if resData == nil {
		err := vendorNode5MinApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "insert failed")
		}
		log.Println("insert succeed")
	}
}
func DeleteVendorNode5MinModel(id string) {
	err := vendorNode5MinApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "delete failed")
	}
	log.Println("delete succeed")
}

func UpdateVendorNode5MinModel(data *hadoop.MeasureCommonUnitList) {
	err := vendorNode5MinApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "update failed")
	}
	log.Println("update succeed")
}

func FindVendorNode5MinModelApiById(id string) *hadoop.MeasureCommonUnitList {
	resData, err := vendorNode5MinApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "query failed")
	}
	log.Println("query succeed")
	return resData
}
