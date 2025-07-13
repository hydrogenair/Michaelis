package logic

import (
	"LingXi/app/star/model"
	"context"

	"LingXi/app/star/cmd/rpc/internal/svc"
	"LingXi/app/star/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfEnterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIfEnterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfEnterLogic {
	return &IfEnterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IfEnterLogic) IfEnter(in *pb.IfEnterRequest) (*pb.IfEnterResponse, error) {
	var userStar *model.UserStar
	userStar, _ = l.svcCtx.UserStarModel.FindOneByUserIdStarId(l.ctx, in.UserId, in.StarId)

	if userStar != nil {
		return &pb.IfEnterResponse{
			Enter: true,
		}, nil
	}
	return &pb.IfEnterResponse{
		Enter: false,
	}, nil
}
