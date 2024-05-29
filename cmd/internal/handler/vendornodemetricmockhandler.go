package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/logic"
	"hadoopmock/cmd/internal/types"
	"net/http"
)

func VendorNodeMetricMockHandler(mgo *config.Mongo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostVendorNodeMetricReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewVendorNodeMetricMockDetailLogic(r.Context())
		resp, err := l.VendorNodeMetricMockDetail(mgo, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
