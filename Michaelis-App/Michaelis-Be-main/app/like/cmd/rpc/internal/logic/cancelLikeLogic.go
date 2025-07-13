package logic

import (
	"LingXi/app/like/cmd/rpc/internal/svc"
	"LingXi/app/like/cmd/rpc/pb"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeLogic {
	return &CancelLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelLikeLogic) CancelLike(in *pb.CancelLikeRequest) (*pb.CancelLikeResponse, error) {
	like, err := l.svcCtx.LikeModel.FindOneByUserIdTypeOutId(l.ctx, in.UserId, in.Type, in.OutId)
	if err != nil {
		return nil, errors.Wrap(err, "haven't find this record")
	}
	if err := l.svcCtx.LikeModel.Delete(l.ctx, nil, like.Id); err != nil {
		return nil, errors.Wrap(err, "failed to delete the like")
	}
	return &pb.CancelLikeResponse{
		Type:   like.Type,
		OutId:  like.OutId,
		UserId: like.UserId,
	}, nil
}
