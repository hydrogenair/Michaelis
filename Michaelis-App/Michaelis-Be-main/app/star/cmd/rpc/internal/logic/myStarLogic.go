package logic

import (
	"context"

	"LingXi/app/star/cmd/rpc/internal/svc"
	"LingXi/app/star/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMyStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyStarLogic {
	return &MyStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MyStarLogic) MyStar(in *pb.MyStarRequest) (*pb.MyStarResponse, error) {
	star, err := l.svcCtx.UserStarModel.FindAll(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	var S = make([]string, len(star))
	for i, m := range star {
		s, _ := l.svcCtx.StarsModel.FindOne(l.ctx, m.StarId)
		S[i] = s.StarName
	}
	return &pb.MyStarResponse{
		Star: S,
	}, nil
}
