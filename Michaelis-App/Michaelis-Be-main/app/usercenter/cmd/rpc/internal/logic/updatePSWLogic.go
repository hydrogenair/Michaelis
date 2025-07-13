package logic

import (
	rds "LingXi/common/redis"
	"LingXi/common/tool"
	"LingXi/common/verifycode"
	"LingXi/common/xerr"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"time"

	"LingXi/app/usercenter/cmd/rpc/internal/svc"
	"LingXi/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrOldPassword = xerr.NewErrMsg("password wrong")

type UpdatePSWLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePSWLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePSWLogic {
	return &UpdatePSWLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePSWLogic) UpdatePSW(in *pb.UpdatePSWReq) (*pb.UpdatePSWResp, error) {
	userId := in.Id
	fmt.Println()
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	if in.Verificationcode != "" {
		if !l.Verify(in.Verificationcode, user.Email) {
			return nil, errors.Wrapf(ErrVerificationError, "The code for verification does not match email:%s", user.Email)
		}
		user.Password = tool.Md5ByString(in.Newpassword)
		user.UpdateTime = time.Now()
		if err := l.svcCtx.UserModel.Update(l.ctx, nil, user); err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Update Info error:%v", err)
		}

	}
	if in.Verificationcode == "" {
		if user.Password != tool.Md5ByString(in.Oldpassword) {
			return nil, errors.Wrap(ErrOldPassword, "Please input correct password")
		}
		user.Password = tool.Md5ByString(in.Newpassword)
		user.UpdateTime = time.Now()
		if err := l.svcCtx.UserModel.Update(l.ctx, nil, user); err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Update Info error:%v", err)
		}
	}

	return &pb.UpdatePSWResp{}, nil
}

func (l *UpdatePSWLogic) Verify(code, email string) bool {
	var store = verifycode.RedisStore{
		RedisClient: &rds.RedisClient{
			Client:  l.svcCtx.RedisClient,
			Context: l.ctx,
		},
		KeyPrefix: l.svcCtx.Config.APP.Name + ":verifycode:",
	}
	return store.Verify(email, code)
}
