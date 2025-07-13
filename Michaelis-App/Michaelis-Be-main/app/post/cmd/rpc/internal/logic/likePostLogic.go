package logic

import (
	"LingXi/app/like/cmd/rpc/likecenter"
	"context"
	"github.com/pkg/errors"

	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikePostLogic {
	return &LikePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikePostLogic) LikePost(in *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	ifLike, err := l.svcCtx.LikeCenterRpc.CheckLike(l.ctx, &likecenter.CheckLikeRequest{
		OutId:  in.PostId,
		UserId: in.UserId,
		Type:   "post",
	})
	if err != nil {
		return nil, err
	}
	if ifLike.Like {
		return nil, errors.New("have liked")
	}
	likeResp, err := l.svcCtx.LikeCenterRpc.PostLike(l.ctx, &likecenter.LikeRequest{
		OutId:  in.PostId,
		UserId: in.UserId,
		Type:   "post",
	})
	if err != nil {
		return nil, err
	}
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	post.LikeNum += 1
	if err := l.svcCtx.PostModel.Update(l.ctx, nil, post); err != nil {
		l.svcCtx.LikeCenterRpc.CancelLike(l.ctx, &likecenter.CancelLikeRequest{
			OutId:  in.PostId,
			UserId: in.UserId,
			Type:   "post",
		})
		return nil, err
	}
	return &pb.LikePostResponse{
		PostId: likeResp.OutId,
	}, nil
}
