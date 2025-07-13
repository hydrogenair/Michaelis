package stars

import (
	"LingXi/app/star/cmd/rpc/starcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/star/cmd/api/internal/svc"
	"LingXi/app/star/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnterStarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnterStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnterStarLogic {
	return &EnterStarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnterStarLogic) EnterStar(req *types.EnterRequest) (resp *types.EnterResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	enterResp, err := l.svcCtx.StarCenter.EnterStar(l.ctx, &starcenter.EnterStarRequest{
		StarId: req.StarId,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.EnterResponse{
		UserId:  enterResp.UserId,
		StarNam: enterResp.StarName,
	}, nil
}
