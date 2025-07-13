package user

import (
	"LingXi/app/usercenter/cmd/api/internal/svc"
	"LingXi/app/usercenter/cmd/api/internal/types"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	infoResp, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		UserInfo: types.User{
			Id:       infoResp.User.Id,
			Email:    infoResp.User.Email,
			NickName: infoResp.User.Nickname,
			Gender:   infoResp.User.Gender,
			Age:      infoResp.User.Age,
			Identity: infoResp.User.Identity,
		},
	}, nil
}
