package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/handler"
)

func main() {
	var c rest.RestConf
	c.Host = "localhost"
	c.Port = 8080
	server := rest.MustNewServer(c)
	defer server.Stop()
	mgo := config.ReadMongoConfig()
	handler.RegisterHandlers(server, mgo)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
