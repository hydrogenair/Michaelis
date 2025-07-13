package logic

import (
	"context"
	"github.com/pkg/errors"

	"LingXi/app/images/cmd/rpc/internal/svc"
	"LingXi/app/images/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllImagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllImagesLogic {
	return &GetAllImagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllImagesLogic) GetAllImages(in *pb.GetAllImagesRequest) (*pb.GetAllImagesResponse, error) {
	images, err := l.svcCtx.ImageModel.GetAllImages(l.ctx, in.GetType(), in.GetOutId())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all images")
	}
	if images != nil {
		Images := make([]*pb.Image, len(images))
		for i, m := range images {
			Images[i] = &pb.Image{
				Id:    m.Id,
				Type:  m.Type,
				OutId: m.OutId,
				Url:   m.Url,
			}
		}
		return &pb.GetAllImagesResponse{
			Images: Images,
		}, nil
	}
	return nil, nil

}
