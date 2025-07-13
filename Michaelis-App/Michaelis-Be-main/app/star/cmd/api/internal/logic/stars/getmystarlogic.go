package stars

import (
	"LingXi/app/star/cmd/rpc/starcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/star/cmd/api/internal/svc"
	"LingXi/app/star/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyStarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyStarLogic {
	return &GetMyStarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyStarLogic) GetMyStar(req *types.MyStarsRequest) (resp *types.MyStarsResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	starResp, err := l.svcCtx.StarCenter.MyStar(l.ctx, &starcenter.MyStarRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.MyStarsResponse{
		Star: starResp.Star,
	}, nil
}
