package logic

import (
	"LingXi/app/comment/cmd/rpc/internal/svc"
	"LingXi/app/comment/cmd/rpc/pb"
	"LingXi/app/comment/model"
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCommentLogic {
	return &PostCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostCommentLogic) PostComment(in *pb.CommentRequest) (*pb.CommentResponse, error) {
	var comment = model.Comments{
		Type:        in.Type,
		PublisherId: in.PublisherId,
		Content:     in.Content,
		LikeNum:     0,
		ReplyNum:    0,
		OutId:       in.OutId,
		CreateTime:  time.Now(),
	}

	if err := l.svcCtx.CommentModel.Insert(l.ctx, nil, &comment); err != nil {
		return nil, errors.Wrapf(err, "failed to send a comment,error:%v", err)
	}

	return &pb.CommentResponse{
		Comment: &pb.Comment{
			CommentId:   comment.Id,
			PublisherId: comment.PublisherId,
			Type:        comment.Type,
			Content:     comment.Content,
			LikeNum:     comment.LikeNum,
			ReplyNum:    comment.ReplyNum,
			OutId:       comment.OutId,
		},
	}, nil
}
