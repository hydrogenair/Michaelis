package logic

import (
	"LingXi/app/like/cmd/rpc/internal/svc"
	"LingXi/app/like/cmd/rpc/pb"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikesLogic {
	return &GetLikesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikesLogic) GetLikes(in *pb.GetLikesRequest) (*pb.GetLikesResponse, error) {
	like, err := l.svcCtx.LikeModel.FindByTpOutId(l.ctx, in.Type, in.OutId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get likes")
	}

	return &pb.GetLikesResponse{
		LikeSum: int64(len(like)),
	}, nil
}
