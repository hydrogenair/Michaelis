package posts

import (
	"LingXi/app/post/cmd/rpc/postcenter"
	"context"
	"fmt"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewPostLogic {
	return &ViewPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewPostLogic) ViewPost(req *types.ViewPostRequest) (resp *types.ViewPostResponse, err error) {
	viewResp, err := l.svcCtx.PostCenterRpc.ViewPost(l.ctx, &postcenter.ViewPostRequest{
		PostId: req.PostId,
	})
	fmt.Println(viewResp)
	if err != nil {
		return nil, err
	}
	return &types.ViewPostResponse{
		PostId: viewResp.PostId,
	}, nil
}
