package posts

import (
	"LingXi/app/post/cmd/rpc/postcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIfLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfLikeLogic {
	return &IfLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IfLikeLogic) IfLike(req *types.IfLikeRequest) (resp *types.IfLikeResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	iflikeResp, err := l.svcCtx.PostCenterRpc.IfLike(l.ctx, &postcenter.IfLikeRequest{
		PostId: req.PostId,
		UserId: userId,
	})

	return &types.IfLikeResponse{iflikeResp.Like}, nil
}
