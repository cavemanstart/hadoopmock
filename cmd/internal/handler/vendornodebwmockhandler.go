package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"hadoopmock/cmd/internal/config"
	"hadoopmock/cmd/internal/logic"
	"hadoopmock/cmd/internal/types"
	"net/http"
)

func VendorNodeBwMockHandler(mgo *config.Mongo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostVendorNodeBwReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewVendorNodeBwDetailLogic(r.Context())
		resp, err := l.VendorNodeBwMockDetail(mgo, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}