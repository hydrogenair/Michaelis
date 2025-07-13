package logic

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"LingXi/common/tool"
	"LingXi/common/xerr"
	"context"
	"github.com/pkg/errors"

	"LingXi/app/usercenter/cmd/rpc/internal/svc"
	"LingXi/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(ErrVerificationError, "userId : %d", in.Email)
	}

	if !(tool.Md5ByString(in.Password) == user.Password) {
		return nil, xerr.NewErrMsg("Wrong Password")
	}
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId:   user.Id,
		Identity: user.Identity,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", user.Id)
	}
	return &pb.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
