package handler

import (
	"LingXi/app/chat/cmd/api/internal/logic/chat"
	"LingXi/app/chat/cmd/api/internal/svc"
	"LingXi/app/chat/cmd/api/internal/types"
	chat1 "LingXi/common/chat"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SocketReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		chat1.ServeWs(svcCtx.Hub, w, r)
		//go Run(svcCtx, r)
		l := chat.NewChatLogic(r.Context(), svcCtx)
		resp, err := l.Chat(&req)
		response.Response(w, resp, err)

	}
}
