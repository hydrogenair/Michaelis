package user

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"context"

	"LingXi/app/usercenter/cmd/api/internal/svc"
	"LingXi/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIdentityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIdentityLogic {
	return &UpdateIdentityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateIdentityLogic) UpdateIdentity(req *types.UpdateIdentityReq) (*types.UpdateIdentityResp, error) {
	AuthResp, err := l.svcCtx.UserCenterRpc.GiveUserAuth(l.ctx, &usercenter.AuthReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateIdentityResp{
		User: types.User{
			Id:       AuthResp.User.Id,
			Email:    AuthResp.User.Email,
			NickName: AuthResp.User.Nickname,
			Gender:   AuthResp.User.Gender,
			Age:      AuthResp.User.Age,
			Identity: AuthResp.User.Identity,
		}}, nil
}
