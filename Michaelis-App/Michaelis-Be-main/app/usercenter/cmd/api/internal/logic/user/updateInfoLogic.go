package user

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/usercenter/cmd/api/internal/svc"
	"LingXi/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateInfoReq) (resp *types.UpdateInfoResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	updateinfoResp, err := l.svcCtx.UserCenterRpc.UpdateInfo(l.ctx, &usercenter.UpdateInfoReq{
		Id:       userId,
		NickName: req.NickName,
		Gender:   req.Gender,
		Age:      req.Age,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateInfoResp{
		UserInfo: types.User{
			Id:       updateinfoResp.User.Id,
			Email:    updateinfoResp.User.Email,
			NickName: updateinfoResp.User.Nickname,
			Gender:   updateinfoResp.User.Gender,
			Age:      updateinfoResp.User.Age,
			Identity: updateinfoResp.User.Identity,
		},
	}, nil
}
