package logic

import (
	"context"

	"LingXi/app/star/cmd/rpc/internal/svc"
	"LingXi/app/star/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type StarDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStarDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StarDetailLogic {
	return &StarDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StarDetailLogic) StarDetail(in *pb.StarDetailRequest) (*pb.StarDetailResponse, error) {
	star, err := l.svcCtx.StarsModel.FindOne(l.ctx, in.StarId)
	if err != nil {
		return nil, err
	}
	return &pb.StarDetailResponse{
		Name:   star.StarName,
		People: star.PeopleCount,
	}, nil
}
