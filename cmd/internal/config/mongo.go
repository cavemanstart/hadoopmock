package config

import (
	"github.com/zeromicro/go-zero/core/conf"
)

type Mongo struct {
	MongoUrl      string
	MongoDatabase string
}

func ReadMongoConfig() *Mongo {
	var mgo Mongo
	conf.MustLoad("cmd/internal/config.yaml", &mgo)
	return &mgo
}
