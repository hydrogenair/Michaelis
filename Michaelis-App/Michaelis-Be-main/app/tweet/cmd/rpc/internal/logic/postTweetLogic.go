package logic

import (
	"LingXi/app/tweet/model"
	"LingXi/common/xerr"
	"context"
	"github.com/pkg/errors"

	"LingXi/app/tweet/cmd/rpc/internal/svc"
	"LingXi/app/tweet/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostTweetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostTweetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostTweetLogic {
	return &PostTweetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostTweetLogic) PostTweet(in *pb.TweetRequest) (*pb.TweetResponse, error) {
	userId := in.Id
	tweet := model.Tweet{
		PublisherId: userId,
		Content:     in.Content,
	}
	if err := l.svcCtx.TweetModel.Insert(l.ctx, nil, &tweet); err != nil {
		return nil, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "failed to publish a tweet.sql")
	}
	return &pb.TweetResponse{
		Tweet: &pb.Tweet{
			Id:          tweet.Id,
			PublisherId: tweet.PublisherId,
			Content:     tweet.Content,
		},
	}, nil
}
