package main

import (
	"fmt"
	"hadoopmock/cmd/internal/mock"
)

func main() {
	//mongoConfig := config.ReadMongoConfig()
	//bwModel := model.NewVendorNodeBwModel(mongoConfig.MongoUrl, mongoConfig.MongoDatabase)
	//_, err := bwModel.FindById(context.Background(), "test")
	//if err != nil {
	//	return
	//}
	//dateString := "2022-01-02"
	//layout := "2006-01-02"
	//
	//tm, err := time.ParseInLocation(layout, dateString, time.Local)
	//if err != nil {
	//	fmt.Println("日期解析出错:", err)
	//	return
	//}
	//fmt.Println(tm.Unix())
	//// 增加一天
	//tm = tm.AddDate(0, 0, 1)
	//
	//fmt.Println("增加一天后的日期:", tm.Format("2006-01-02"))
	nodeMetric := mock.MockVendorNodeMetric("2021-01-12", "2021-01-28", 9)
	fmt.Println("nodeMetric:", nodeMetric)
}
