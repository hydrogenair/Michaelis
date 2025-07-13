package user

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"LingXi/common/xerr"
	"context"
	"github.com/pkg/errors"

	"LingXi/app/usercenter/cmd/api/internal/svc"
	"LingXi/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.VerificationReq) (resp *types.VerfificationResp, err error) {
	verifyResp, err := l.svcCtx.UserCenterRpc.GetVerification(l.ctx, &usercenter.VerificationReq{
		Email: req.Email,
	})

	if err != nil {

		return nil, errors.Wrap(xerr.NewErrMsg("send verification code fail"), "")
	}
	return &types.VerfificationResp{
		VerficationCode: verifyResp.Verification,
	}, nil
}
