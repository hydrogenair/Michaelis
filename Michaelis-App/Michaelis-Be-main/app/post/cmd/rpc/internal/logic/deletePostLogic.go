package logic

import (
	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"
	"LingXi/common/xerr"
	"context"
	"database/sql"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePostLogic) DeletePost(in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		return nil, xerr.NewErrMsg("Don't have this post")
	}
	if post.PublisherId != in.PublisherId {
		return nil, xerr.NewErrMsg("This is not your post")
	}
	post.DeleteTime = sql.NullTime{Time: time.Now(), Valid: true}
	if err = l.svcCtx.PostModel.Update(l.ctx, nil, post); err != nil {
		return nil, xerr.NewErrMsg("删除失败")
	}
	return &pb.DeletePostResponse{
		PostId: post.Id,
	}, nil
}
