package logic

import (
	"LingXi/app/comment/cmd/rpc/internal/svc"
	"LingXi/app/comment/cmd/rpc/pb"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	err := l.svcCtx.CommentModel.Delete(l.ctx, nil, in.CommentId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete")
	}
	return &pb.DeleteCommentResponse{
		CommentId: in.CommentId,
	}, nil
}
