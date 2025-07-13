package image

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"context"
	"fmt"

	"LingXi/app/images/cmd/api/internal/svc"
	"LingXi/app/images/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllImagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllImagesLogic {
	return &GetAllImagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllImagesLogic) GetAllImages(req *types.GetAllImagesRequest) (resp *types.GetAllImagesResponse, err error) {
	fmt.Println(req)
	imageResp, err := l.svcCtx.ImageCenter.GetAllImages(l.ctx, &imagecenter.GetAllImagesRequest{
		Type:  req.Type,
		OutId: req.OutId,
	})
	resp = &types.GetAllImagesResponse{}
	images := make([]types.Image, len(imageResp.Images))
	for i, m := range imageResp.Images {
		images[i] = types.Image{
			Id:    m.GetId(),
			Type:  m.GetType(),
			OutId: m.GetOutId(),
			Url:   m.GetUrl(),
		}
	}
	resp.Images = images
	return resp, nil
}
