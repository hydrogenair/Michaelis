package handler

import (
	"LingXi/app/tweet/cmd/api/internal/logic/tweeet"
	"LingXi/app/tweet/cmd/api/internal/svc"
	"LingXi/app/tweet/cmd/api/internal/types"
	"github.com/Wishforpeace/zero-tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func PostTweetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TweetRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tweeet.NewPostTweetLogic(r.Context(), svcCtx)
		resp, err := l.PostTweet(&req)
		response.Response(w, resp, err)

	}
}
