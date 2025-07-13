package logic

import (
	"context"
	"fmt"

	"LingXi/app/usercenter/cmd/rpc/internal/svc"
	"LingXi/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	userId := in.Id
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	fmt.Println(userInfo.Identity)
	return &pb.UserInfoResp{
		User: &pb.User{
			Id:       userInfo.Id,
			Email:    userInfo.Email,
			Nickname: userInfo.Nickname,
			Gender:   userInfo.Gender,
			Age:      userInfo.Age,
			Identity: userInfo.Identity,
		},
	}, nil
}
