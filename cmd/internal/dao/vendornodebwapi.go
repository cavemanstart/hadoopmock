package dao

import (
	"context"
	"log"

	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
)

var (
	vendorNodeBwApiModel model.VendorNodeBwApiModel
)

func init() {
	mongoConfig := config.ReadConfig()
	vendorNodeBwApiModel = model.NewVendorNodeBwApiModel(mongoConfig.Mongo.Url, mongoConfig.Mongo.Database)
}
func InsertVendorNodeBwModel(data *hadoop.VendorNodeMomentData) {
	resData := FindVendorNodeBwModelApiById(data.Id)
	if resData == nil {
		err := vendorNodeBwApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "measure api insert failed")
		}
		log.Println("measure api insert succeed")
	}
}
func DeleteVendorNodeBwModel(id string) {
	err := vendorNodeBwApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "measure api delete failed")
	}
	log.Println("measure api delete succeed")
}

func UpdateVendorNodeBwModel(data *hadoop.VendorNodeMomentData) {
	err := vendorNodeBwApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "measure api update failed")
	}
	log.Println("measure api update succeed")
}

func FindVendorNodeBwModelApiById(id string) *hadoop.VendorNodeMomentData {
	resData, err := vendorNodeBwApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, " measure api query failed")
	}
	log.Println(" measure api query succeed")
	return resData
}
