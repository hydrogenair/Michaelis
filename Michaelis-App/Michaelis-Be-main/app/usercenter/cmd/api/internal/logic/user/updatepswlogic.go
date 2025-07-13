package user

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	ctxdata "LingXi/common/cxtdata"
	"context"
	"github.com/jinzhu/copier"

	"LingXi/app/usercenter/cmd/api/internal/svc"
	"LingXi/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePSWLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePSWLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePSWLogic {
	return &UpdatePSWLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePSWLogic) UpdatePSW(req *types.UpdatePSWReq) (*types.UpdatePSWResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	PswResp, err := l.svcCtx.UserCenterRpc.UpdatePSW(l.ctx, &usercenter.UpdatePSWReq{
		Id:               userId,
		Oldpassword:      req.OldPassword,
		Newpassword:      req.NewPassword,
		Verificationcode: req.VerficationCode,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UpdatePSWResp
	_ = copier.Copy(&resp, PswResp)
	return &resp, nil
}
