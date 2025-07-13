package logic

import (
	"context"
	"github.com/pkg/errors"

	"LingXi/app/images/cmd/rpc/internal/svc"
	"LingXi/app/images/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneImageLogic {
	return &GetOneImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneImageLogic) GetOneImage(in *pb.GetOneImageRequest) (*pb.GetOneImageResponse, error) {
	image, err := l.svcCtx.ImageModel.GetImage(l.ctx, in.GetType(), in.GetOutId())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get image")
	}
	return &pb.GetOneImageResponse{
		Image: &pb.Image{
			Id:    image.Id,
			Type:  image.Type,
			OutId: image.OutId,
			Url:   image.Url,
		},
	}, nil
}
