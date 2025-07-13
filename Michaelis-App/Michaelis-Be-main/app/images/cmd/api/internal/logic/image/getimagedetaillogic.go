package image

import (
	"LingXi/app/images/cmd/api/internal/svc"
	"LingXi/app/images/cmd/api/internal/types"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetImageDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetImageDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImageDetailLogic {
	return &GetImageDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetImageDetailLogic) GetImageDetail(req *types.GetImageDetailRequest) (resp *types.GetImageDetailResponse, err error) {
	//fmt.Println(req.Id)
	imageResp, err := l.svcCtx.ImageCenter.GetDetail(l.ctx, &imagecenter.GetDetailRequest{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}
	return &types.GetImageDetailResponse{
		Image: types.Image{
			Id:    imageResp.Images.Id,
			Type:  imageResp.Images.Type,
			OutId: imageResp.Images.OutId,
			Url:   imageResp.Images.Url,
		},
	}, nil
}
