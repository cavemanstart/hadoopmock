package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type MongoConfig struct {
	Mongo struct {
		Url      string `yaml:"url"`
		Database string `yaml:"database"`
	} `yaml:"mongo"`
}

func ReadConfig() *MongoConfig {
	data, err := os.ReadFile("../config.yaml")
	if err != nil {
		return nil
	}
	var mongoConfig MongoConfig
	err = yaml.Unmarshal(data, &mongoConfig)
	fmt.Println("url", mongoConfig.Mongo.Url)
	if err != nil {
		return nil
	}
	return &mongoConfig
}
