package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"hadoopmock/cmd/internal/config"
	"net/http"
	"time"
)

func RegisterHandlers(server *rest.Server, mgo *config.Mongo) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/node/metric",
					Handler: VendorNodeMetricMockHandler(mgo),
				},
			}...,
		),
		rest.WithPrefix("/billing"),
		rest.WithTimeout(20000*time.Millisecond),
	)
}
