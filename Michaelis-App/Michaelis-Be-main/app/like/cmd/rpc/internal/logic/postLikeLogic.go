package logic

import (
	"LingXi/app/like/cmd/rpc/internal/svc"
	"LingXi/app/like/cmd/rpc/pb"
	"LingXi/app/like/model"
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLikeLogic {
	return &PostLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostLikeLogic) PostLike(in *pb.LikeRequest) (*pb.LikeResponse, error) {
	var like = model.Like{
		UserId:     in.UserId,
		Type:       in.Type,
		OutId:      in.OutId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := l.svcCtx.LikeModel.Insert(l.ctx, nil, &like); err != nil {
		return nil, errors.Wrap(err, "failed to insert a like")
	}

	return &pb.LikeResponse{
		UserId: like.UserId,
		Type:   like.Type,
		OutId:  like.OutId,
	}, nil
}
