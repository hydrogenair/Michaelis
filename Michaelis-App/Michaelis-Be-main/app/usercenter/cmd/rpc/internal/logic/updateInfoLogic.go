package logic

import (
	"LingXi/app/usercenter/model"
	"LingXi/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"LingXi/app/usercenter/cmd/rpc/internal/svc"
	"LingXi/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInfoLogic) UpdateInfo(in *pb.UpdateInfoReq) (*pb.UpdateInfoResp, error) {
	userId := in.Id
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	var user model.User
	user = *userInfo

	//if userInfo.Age != in.Age {
	//	user.Age = in.Age
	//}
	//
	//if userInfo.Gender != in.Gender {
	//	user.Gender = in.Gender
	//}
	//if user.Nickname != in.NickName {
	//	user.Nickname = in.NickName
	//}
	switch {
	case userInfo.Age != in.Age:
		user.Age = in.Age
	case userInfo.Gender != in.Gender:
		user.Gender = in.Gender
	case user.Nickname != in.NickName:
		user.Nickname = in.NickName
	default:
		return nil, err
	}

	user.UpdateTime = time.Now()
	if err := l.svcCtx.UserModel.Update(l.ctx, nil, &user); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Update Info error:%v", err)
	}

	userInfo, err = l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Get Info error:%v", err)
	}
	return &pb.UpdateInfoResp{
		User: &pb.User{
			Id:       userInfo.Id,
			Email:    userInfo.Email,
			Nickname: userInfo.Nickname,
			Gender:   userInfo.Gender,
			Age:      userInfo.Age,
		},
	}, nil
}
