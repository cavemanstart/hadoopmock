package mock

import (
	"context"
	"log"
	"time"

	"hadoopmock/cmd/internal/hadoop"
	"hadoopmock/cmd/internal/model"
)

var (
	mongoUrl               = "mongodb://localhost:27017"
	mongoDb         string = "hadoop"
	measureApiModel model.MeasureApiModel
	data            = map[string]hadoop.MeasureCommonUnit{
		"2023-06-01": {Time: time.Now().Unix(), Bandwidth: 0},
		"2023-06-02": {Time: time.Now().Unix(), Bandwidth: 0},
		"2023-06-03": {Time: time.Now().Unix(), Bandwidth: 0},
		"2023-06-04": {Time: time.Now().Unix(), Bandwidth: 0},
		"2023-06-05": {Time: time.Now().Unix(), Bandwidth: 400000},
		"2023-06-06": {Time: time.Now().Unix(), Bandwidth: 500000},
		"2023-06-07": {Time: time.Now().Unix(), Bandwidth: 600000},
		"2023-06-08": {Time: time.Now().Unix(), Bandwidth: 700000},
		"2023-06-09": {Time: time.Now().Unix(), Bandwidth: 0},
	}
)

func InsertMeasureApi(id string) {
	if measureApiModel == nil {
		measureApiModel = model.NewMeasureApiModel(mongoUrl, mongoDb)
	}
	mockMeasureData := hadoop.MeasureCommonData{Id: id, DayPeak: data, Peak: hadoop.MeasureCommonUnit{Bandwidth: 5000}}
	err := measureApiModel.Insert(context.Background(), &mockMeasureData)
	if err != nil {
		log.Fatal(err, "measure api insert failed")
	}
	log.Println("measure api insert succeed")
}
func DeleteMeasureApi(id string) {
	if measureApiModel == nil {
		measureApiModel = model.NewMeasureApiModel(mongoUrl, mongoDb)
	}
	err := measureApiModel.DeleteById(context.Background(), id)
	if err != nil {
		log.Fatal(err, "measure api delete failed")
	}
	log.Println("measure api delete succeed")
}

func UpdateMeasureApi(id string) {
	if measureApiModel == nil {
		measureApiModel = model.NewMeasureApiModel(mongoUrl, mongoDb)
	}
	mockMeasureData := hadoop.MeasureCommonData{Id: id, DayPeak: data, Peak: hadoop.MeasureCommonUnit{Bandwidth: 2000}}
	err := measureApiModel.Update(context.Background(), &mockMeasureData)
	if err != nil {
		log.Fatal(err, "measure api update failed")
	}
	log.Println("measure api update succeed")
}

func FindMeasureApiById(id string) *hadoop.MeasureCommonData {
	if measureApiModel == nil {
		measureApiModel = model.NewMeasureApiModel(mongoUrl, mongoDb)
	}
	resData, err := measureApiModel.FindById(context.Background(), id)
	if err != nil {
		log.Fatal(err, " measure api query failed")
	}
	log.Println(" measure api query succeed")
	return resData
}
