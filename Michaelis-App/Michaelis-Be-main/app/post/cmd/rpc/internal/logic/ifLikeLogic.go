package logic

import (
	"LingXi/app/like/cmd/rpc/likecenter"
	"context"

	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIfLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfLikeLogic {
	return &IfLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IfLikeLogic) IfLike(in *pb.IfLikeRequest) (*pb.IfLikeResponse, error) {
	like, err := l.svcCtx.LikeCenterRpc.CheckLike(l.ctx, &likecenter.CheckLikeRequest{
		OutId:  in.PostId,
		UserId: in.UserId,
		Type:   "post",
	})
	if err != nil {
		return nil, err
	}
	return &pb.IfLikeResponse{
		Like: like.Like,
	}, nil
}
