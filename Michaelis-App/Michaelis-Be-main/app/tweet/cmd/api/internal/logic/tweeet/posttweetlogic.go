package tweeet

import (
	"LingXi/app/tweet/cmd/rpc/tweetcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/tweet/cmd/api/internal/svc"
	"LingXi/app/tweet/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostTweetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostTweetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostTweetLogic {
	return &PostTweetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostTweetLogic) PostTweet(req *types.TweetRequest) (resp *types.TweetResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	tweetResp, err := l.svcCtx.TweetCenterRpc.PostTweet(l.ctx, &tweetcenter.TweetRequest{
		Id:      userId,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &types.TweetResponse{
		Tweet: types.Tweet{
			Id:          tweetResp.Tweet.Id,
			PublisherId: tweetResp.Tweet.PublisherId,
			Content:     tweetResp.Tweet.Content,
		},
	}, nil
}
