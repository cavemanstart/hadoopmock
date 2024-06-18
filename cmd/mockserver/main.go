package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/rest"

	"hadoopmock/cmd/mockserver/config"
	"hadoopmock/cmd/mockserver/handler"
)

var configFile = flag.String("f", "cmd/mockserver/config.yaml", "the config file")

func main() {
	flag.Parse()
	cfg := config.ReadConfig(*configFile)
	var c rest.RestConf
	c.Host = cfg.Host
	c.Port = cfg.Port
	server := rest.MustNewServer(c)
	defer server.Stop()

	sc := config.NewServiceContext(cfg)
	handler.RegisterHandlers(server, sc)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
