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
func InsertVendorNodeBwModel(data *hadoop.NodeMomentDataList) {
	resData := FindVendorNodeBwModelApiById(data.Id)
	if resData == nil {
		err := vendorNodeBwApiModel.Insert(context.Background(), data)
		if err != nil {
			log.Fatal(err, "insert failed")
		}
		log.Println("insert succeed")
	}
}
func DeleteVendorNodeBwModel(id string) {
	err := vendorNodeBwApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "delete failed")
	}
	log.Println("delete succeed")
}

func UpdateVendorNodeBwModel(data *hadoop.NodeMomentDataList) {
	err := vendorNodeBwApiModel.Update(context.Background(), data)
	if err != nil {
		log.Fatal(err, "update failed")
	}
	log.Println("update succeed")
}

func FindVendorNodeBwModelApiById(id string) *hadoop.NodeMomentDataList {
	resData, err := vendorNodeBwApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, " query failed")
	}
	log.Println(" query succeed")
	return resData
}
