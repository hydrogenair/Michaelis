package posts

import (
	"LingXi/app/post/cmd/rpc/postcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikePostLogic {
	return &LikePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikePostLogic) LikePost(req *types.LikePostRequest) (resp *types.LikePostResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	likeResp, err := l.svcCtx.PostCenterRpc.LikePost(l.ctx, &postcenter.LikePostRequest{
		PostId: req.PostId,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.LikePostResponse{
		PostId: likeResp.PostId,
	}, nil
}
