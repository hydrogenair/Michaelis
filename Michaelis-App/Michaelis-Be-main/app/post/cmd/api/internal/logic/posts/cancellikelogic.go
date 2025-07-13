package posts

import (
	"LingXi/app/post/cmd/rpc/postcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeLogic {
	return &CancelLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLikeLogic) CancelLike(req *types.CancelLikeRequest) (resp *types.CancelLikeResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	cancelResp, err := l.svcCtx.PostCenterRpc.CancelLike(l.ctx, &postcenter.CancelRequest{
		PostId: req.PostId,
		UserId: userId,
	})
	return &types.CancelLikeResponse{
		PostId: cancelResp.PostId,
	}, nil
}
