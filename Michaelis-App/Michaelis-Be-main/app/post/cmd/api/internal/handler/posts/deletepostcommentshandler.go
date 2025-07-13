package handler

import (
	"LingXi/app/post/cmd/api/internal/logic/posts"
	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func DeletePostCommentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCommentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := posts.NewDeletePostCommentsLogic(r.Context(), svcCtx)
		resp, err := l.DeletePostComments(&req)
		response.Response(w, resp, err)

	}
}
