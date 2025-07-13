package logic

import (
	"LingXi/app/images/model"
	"context"
	"github.com/pkg/errors"
	"time"

	"LingXi/app/images/cmd/rpc/internal/svc"
	"LingXi/app/images/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostImageLogic {
	return &PostImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostImageLogic) PostImage(in *pb.PostImageRequest) (*pb.PostImageResponse, error) {
	var image = model.Images{
		//Id:         ,
		Type:       in.Type,
		OutId:      in.OutId,
		Url:        in.Url,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: sql.NullTime{},
	}

	if err := l.svcCtx.ImageModel.Insert(l.ctx, nil, &image); err != nil {
		return nil, errors.Wrap(err, "failed to insert an image")
	}

	return &pb.PostImageResponse{
		Image: &pb.Image{
			Id:    image.Id,
			Type:  image.Type,
			OutId: image.OutId,
			Url:   image.Url,
		},
	}, nil
}
