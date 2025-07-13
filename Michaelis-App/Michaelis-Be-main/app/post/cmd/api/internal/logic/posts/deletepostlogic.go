package posts

import (
	"LingXi/app/post/cmd/rpc/postcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostRequest) (resp *types.DeletePostResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	deleteResp, err := l.svcCtx.PostCenterRpc.DeletePost(l.ctx, &postcenter.DeletePostRequest{
		PublisherId: userId,
		PostId:      req.PostId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.DeletePostResponse{
		PostId: deleteResp.PostId,
	}
	return resp, nil
}
