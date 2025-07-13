package stars

import (
	"LingXi/app/star/cmd/rpc/starcenter"
	"context"

	"LingXi/app/star/cmd/api/internal/svc"
	"LingXi/app/star/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStarDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStarDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStarDetailLogic {
	return &GetStarDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStarDetailLogic) GetStarDetail(req *types.StarDetailRequest) (resp *types.StarDetailResponse, err error) {
	starDetail, err := l.svcCtx.StarCenter.StarDetail(l.ctx, &starcenter.StarDetailRequest{
		StarId: req.StarId,
	})
	if err != nil {
		return nil, err
	}
	return &types.StarDetailResponse{
		Name:        starDetail.Name,
		PeopleCount: int(starDetail.People),
	}, nil
}
