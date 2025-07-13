package logic

import (
	"LingXi/app/star/model"
	"context"
	"github.com/pkg/errors"

	"LingXi/app/star/cmd/rpc/internal/svc"
	"LingXi/app/star/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnterStarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnterStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnterStarLogic {
	return &EnterStarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EnterStarLogic) EnterStar(in *pb.EnterStarRequest) (*pb.EnterStarResponse, error) {
	userStar, err := l.svcCtx.UserStarModel.FindOneByUserIdStarId(l.ctx, in.GetUserId(), in.GetStarId())
	//if err != nil {
	//	return nil, err
	//}
	if userStar != nil {
		return nil, errors.New("have entered this star")
	}
	var userStar1 = model.UserStar{
		UserId: in.UserId,
		StarId: in.StarId,
	}
	if err := l.svcCtx.UserStarModel.Insert(l.ctx, nil, &userStar1); err != nil {
		return nil, err
	}
	star, err := l.svcCtx.StarsModel.FindOne(l.ctx, in.GetStarId())
	if err != nil {
		return nil, err
	}
	return &pb.EnterStarResponse{
		UserId:   userStar1.UserId,
		StarName: star.StarName,
	}, nil
}
