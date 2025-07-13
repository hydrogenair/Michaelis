package user

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"LingXi/app/usercenter/cmd/api/internal/svc"
	"LingXi/app/usercenter/cmd/api/internal/types"
	_ "github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	registerResp, err := l.svcCtx.UserCenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Email:        req.Email,
		Password:     req.Password,
		Verification: req.VerificationCode,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)
	return &resp, nil
}
