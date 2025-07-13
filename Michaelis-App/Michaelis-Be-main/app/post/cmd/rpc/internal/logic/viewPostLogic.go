package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"

	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewPostLogic {
	return &ViewPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewPostLogic) ViewPost(in *pb.ViewPostRequest) (*pb.ViewPostResponse, error) {
	fmt.Println("-------------")
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	fmt.Println(post)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find this post")
	}
	post.ViewNum += 1
	if err := l.svcCtx.PostModel.Update(l.ctx, nil, post); err != nil {
		return nil, errors.Wrap(err, "failed to update the view num")
	}

	return &pb.ViewPostResponse{
		PostId: post.Id,
	}, nil
}
