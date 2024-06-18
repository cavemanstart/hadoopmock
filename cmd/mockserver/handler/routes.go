package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"hadoopmock/cmd/mockserver/config"
	"hadoopmock/cmd/mockserver/util"
)

func RegisterHandlers(server *rest.Server, sc *config.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/node/metric",
					Handler: VendorNodeMetricMockHandler(sc),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/node/bandwidth",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/node/metric",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/node/bandwidth",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/node/5min",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/node/5min",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/evening/metric",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/evening/metric",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/vendor/pernode/metric",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/customer/pernode/metric",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/line/bandwidth/series",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/line/moment/bandwidth",
					Handler: emptyResponse,
				},
				{
					Method:  http.MethodPost,
					Path:    "/v2/pingloss/node/rate",
					Handler: emptyResponse,
				},
			}...,
		),
		rest.WithPrefix("/billing"),
		rest.WithTimeout(20000*time.Millisecond),
	)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v2/bill/mock/measure",
					Handler: GenerateMeasureDataHandler(sc),
				},
			}...,
		),
		rest.WithPrefix("/billing"),
		rest.WithTimeout(20000*time.Millisecond),
	)
}

func emptyResponse(w http.ResponseWriter, r *http.Request) {
	httpx.OkJsonCtx(r.Context(), w, util.SuccessResp([]string{}))
}
