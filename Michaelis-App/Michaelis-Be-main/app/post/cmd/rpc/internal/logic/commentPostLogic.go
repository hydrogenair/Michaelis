package logic

import (
	"context"
	"github.com/pkg/errors"

	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPostLogic {
	return &CommentPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentPostLogic) CommentPost(in *pb.CommentPostRequest) (*pb.CommentPostResponse, error) {
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		return nil, errors.Wrap(err, "database error")
	}
	post.CommentNum += 1
	if err = l.svcCtx.PostModel.Update(l.ctx, nil, post); err != nil {
		return nil, errors.Wrapf(err, "update post comment num error:%v", err.Error())
	}
	return &pb.CommentPostResponse{
		PostId: post.Id,
	}, nil
}
