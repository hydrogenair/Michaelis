package handler

import (
	"LingXi/app/star/cmd/api/internal/logic/stars"
	"LingXi/app/star/cmd/api/internal/svc"
	"LingXi/app/star/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func IfEnterStarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IfEnterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := stars.NewIfEnterStarLogic(r.Context(), svcCtx)
		resp, err := l.IfEnterStar(&req)
		response.Response(w, resp, err)

	}
}
