package handler

import (
	"LingXi/app/images/cmd/api/internal/logic/image"
	"LingXi/app/images/cmd/api/internal/svc"
	"LingXi/app/images/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetAllImagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllImagesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := image.NewGetAllImagesLogic(r.Context(), svcCtx)
		resp, err := l.GetAllImages(&req)
		response.Response(w, resp, err)

	}
}
