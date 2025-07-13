package logic

import (
	"LingXi/app/post/model"
	"LingXi/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPostLogic {
	return &SendPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPostLogic) SendPost(in *pb.PostRequest) (*pb.PostResponse, error) {
	var post = model.Post{
		//Id:          0,
		PublisherId: in.PublisherId,
		Content:     in.Content,
		Category:    in.Category,
		ViewNum:     0,
		LikeNum:     0,
		CommentNum:  0,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		//DeleteTime:  sql.NullTime{},
	}

	if err := l.svcCtx.PostModel.Insert(l.ctx, nil, &post); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Post db user Insert err:%v", err)
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      post.Id,
			PublisherId: post.PublisherId,
			Category:    post.Category,
			Content:     post.Content,
			ViewNum:     post.ViewNum,
			LikeNum:     post.LikeNum,
			CommentNum:  post.CommentNum,
			Time:        post.CreateTime.String(),
		},
	}, nil
}
