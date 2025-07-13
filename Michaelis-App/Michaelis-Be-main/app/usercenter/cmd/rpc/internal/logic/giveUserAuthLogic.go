package logic

import (
	"context"
	"time"

	"LingXi/app/usercenter/cmd/rpc/internal/svc"
	"LingXi/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiveUserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGiveUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveUserAuthLogic {
	return &GiveUserAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GiveUserAuthLogic) GiveUserAuth(in *pb.AuthReq) (*pb.AuthResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	user.Identity = 1
	user.UpdateTime = time.Now()
	user = nil
	user, err = l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.AuthResp{
		User: &pb.User{
			Id:       user.Id,
			Email:    user.Email,
			Nickname: user.Nickname,
			Gender:   user.Gender,
			Age:      user.Age,
			Identity: user.Identity,
		},
	}, nil
}
