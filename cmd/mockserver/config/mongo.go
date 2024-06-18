package config

import (
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	MongoUrl      string
	MongoDatabase string
	Host          string
	Port          int
}

func ReadConfig(path string) *Config {
	var cfg Config
	conf.MustLoad(path, &cfg)
	return &cfg
}
