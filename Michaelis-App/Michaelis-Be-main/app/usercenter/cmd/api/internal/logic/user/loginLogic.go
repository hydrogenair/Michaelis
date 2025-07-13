package user

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"context"

	"LingXi/app/usercenter/cmd/api/internal/svc"
	"LingXi/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.UserCenterRpc.Login(l.ctx, &usercenter.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  loginResp.AccessToken,
		AccessExpire: loginResp.AccessExpire,
		RefreshAfter: loginResp.RefreshAfter,
	}, nil
}
