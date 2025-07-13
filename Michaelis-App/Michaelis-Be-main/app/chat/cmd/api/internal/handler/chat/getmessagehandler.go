package handler

import (
	"LingXi/app/chat/cmd/api/internal/logic/chat"
	"LingXi/app/chat/cmd/api/internal/svc"
	"LingXi/app/chat/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := chat.NewGetMessageLogic(r.Context(), svcCtx)
		resp, err := l.GetMessage(&req)
		response.Response(w, resp, err)
	}
}
