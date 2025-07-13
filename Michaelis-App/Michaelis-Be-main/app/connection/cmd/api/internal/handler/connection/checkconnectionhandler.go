package handler

import (
	"LingXi/app/connection/cmd/api/internal/logic/connection"
	"LingXi/app/connection/cmd/api/internal/svc"
	"LingXi/app/connection/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CheckConnectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetConnectionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := connection.NewCheckConnectionLogic(r.Context(), svcCtx)
		resp, err := l.CheckConnection(&req)
		response.Response(w, resp, err)

	}
}
