package handler

import (
	"LingXi/app/connection/cmd/api/internal/logic/connection"
	"LingXi/app/connection/cmd/api/internal/svc"
	"LingXi/app/connection/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AddConnectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddConnectionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := connection.NewAddConnectionLogic(r.Context(), svcCtx)
		resp, err := l.AddConnection(&req)
		response.Response(w, resp, err)

	}
}
