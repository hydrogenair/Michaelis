package handler

import (
	"LingXi/app/post/cmd/api/internal/logic/posts"
	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetPersonalCategoryPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PersonalPostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := posts.NewGetPersonalCategoryPostLogic(r.Context(), svcCtx)
		resp, err := l.GetPersonalCategoryPost(&req)
		response.Response(w, resp, err)

	}
}
