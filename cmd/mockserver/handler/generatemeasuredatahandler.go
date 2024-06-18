package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"hadoopmock/cmd/mockserver/config"
	"hadoopmock/cmd/mockserver/logic"
	"hadoopmock/cmd/mockserver/types"
	"net/http"
)

func GenerateMeasureDataHandler(sc *config.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostGenerateMeasureReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewGenerateMeasureDetailLogic(r.Context())
		err := l.GenerateMeasureDetail(sc, &req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}
