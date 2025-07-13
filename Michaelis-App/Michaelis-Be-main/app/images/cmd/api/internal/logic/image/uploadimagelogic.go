package image

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"context"

	"LingXi/app/images/cmd/api/internal/svc"
	"LingXi/app/images/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(req *types.UploadImageRequest) (resp *types.UploadImageResponse, err error) {
	imageResp, err := l.svcCtx.ImageCenter.PostImage(l.ctx, &imagecenter.PostImageRequest{
		Type:  req.Type,
		OutId: req.OutId,
		Url:   req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &types.UploadImageResponse{
		Image: types.Image{
			Id:    imageResp.Image.Id,
			Type:  imageResp.Image.Type,
			OutId: imageResp.Image.OutId,
			Url:   imageResp.Image.Url,
		},
	}, nil

}
