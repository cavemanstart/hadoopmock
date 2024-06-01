package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/logic"
	"hadoopmock/cmd/internal/types"
	"net/http"
)

func VendorEveningMetricMockHandler(mgo *config.Mongo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostVendorEveningMetricReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewVendorEveningMetricDetailLogic(r.Context())
		resp, err := l.VendorEveningMetricMockDetail(mgo, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
