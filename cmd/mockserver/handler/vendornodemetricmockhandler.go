package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"hadoopmock/cmd/mockserver/config"
	"hadoopmock/cmd/mockserver/logic"
	"hadoopmock/cmd/mockserver/types"
	"hadoopmock/cmd/mockserver/util"
)

func VendorNodeMetricMockHandler(sc *config.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostVendorNodeMetricReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewVendorNodeMetricDetailLogic(r.Context())
		data, err := l.VendorNodeMetricMockDetail(sc, &req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, util.ErrResp())
		} else {
			httpx.OkJsonCtx(r.Context(), w, util.SuccessResp(data))
		}
	}
}
