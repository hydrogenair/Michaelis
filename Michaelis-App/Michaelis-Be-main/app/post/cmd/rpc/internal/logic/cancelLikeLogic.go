package logic

import (
	"LingXi/app/like/cmd/rpc/likecenter"
	"context"

	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"

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

func (l *CancelLikeLogic) CancelLike(in *pb.CancelRequest) (*pb.CancelResponse, error) {
	cancelLikeResp, err := l.svcCtx.LikeCenterRpc.CancelLike(l.ctx, &likecenter.CancelLikeRequest{
		OutId:  in.PostId,
		UserId: in.UserId,
		Type:   "post",
	})
	if err != nil {
		return nil, err
	}
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	post.LikeNum -= 1
	if err := l.svcCtx.PostModel.Update(l.ctx, nil, post); err != nil {
		l.svcCtx.LikeCenterRpc.PostLike(l.ctx, &likecenter.LikeRequest{
			OutId:  in.PostId,
			UserId: in.UserId,
			Type:   "post",
		})
		return nil, err
	}
	return &pb.CancelResponse{
		PostId: cancelLikeResp.OutId,
	}, nil
}
