package logic

import (
	"LingXi/app/like/cmd/rpc/internal/svc"
	"LingXi/app/like/cmd/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLikeLogic {
	return &CheckLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLikeLogic) CheckLike(in *pb.CheckLikeRequest) (*pb.CheckLikeResponse, error) {
	like, err := l.svcCtx.LikeModel.FindOneByUserIdTypeOutId(l.ctx, in.UserId, in.Type, in.OutId)
	if err != nil {
		return &pb.CheckLikeResponse{
			Like: false,
		}, nil
	}
	if like == nil {
		return &pb.CheckLikeResponse{
			Like: false,
		}, nil
	}
	return &pb.CheckLikeResponse{
		Like: true,
	}, nil
}
