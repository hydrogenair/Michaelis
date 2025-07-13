package stars

import (
	"LingXi/app/star/cmd/rpc/starcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/star/cmd/api/internal/svc"
	"LingXi/app/star/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfEnterStarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIfEnterStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfEnterStarLogic {
	return &IfEnterStarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IfEnterStarLogic) IfEnterStar(req *types.IfEnterRequest) (resp *types.IfEnterResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	ifEnterResp, err := l.svcCtx.StarCenter.IfEnter(l.ctx, &starcenter.IfEnterRequest{
		StarId: req.StarId,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.IfEnterResponse{Enter: ifEnterResp.Enter}, nil
}
