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
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/node/bandwidth",
					Handler: VendorNodeBwMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/node/metric",
					Handler: CustomerNodeMetricMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/node/bandwidth",
					Handler: CustomerNodeBwMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/node/5min",
					Handler: VendorNode5MinMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/node/5min",
					Handler: CustomerNode5MinMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/evening/metric",
					Handler: VendorEveningMetricMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/evening/metric",
					Handler: CustomerEveningMetricMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/pernode/metric",
					Handler: VendorPerNodeMetricMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/pernode/metric",
					Handler: CustomerPerNodeMetricMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/line/bandwidth/series",
					Handler: LineBwSeriesMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/line/moment/bandwidth",
					Handler: LineMomentBwMockHandler(mgo),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/pingloss/node/rate",
					Handler: PingLossNodeRateMockHandler(mgo),
				},
			}...,
		),
		//rest.WithPrefix("/billing"),
		rest.WithTimeout(20000*time.Millisecond),
	)
}
