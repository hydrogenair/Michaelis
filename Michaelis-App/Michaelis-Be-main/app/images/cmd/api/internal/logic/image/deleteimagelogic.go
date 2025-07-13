package image

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"context"

	"LingXi/app/images/cmd/api/internal/svc"
	"LingXi/app/images/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteImageLogic {
	return &DeleteImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteImageLogic) DeleteImage(req *types.DeleteImageRequest) (resp *types.DeleteImageReponse, err error) {
	deleteResp, err := l.svcCtx.ImageCenter.DeleteImage(l.ctx, &imagecenter.DeleteImageRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteImageReponse{
		Id: deleteResp.Id,
	}, nil
}
