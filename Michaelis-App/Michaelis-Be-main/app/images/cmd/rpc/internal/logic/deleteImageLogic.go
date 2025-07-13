package logic

import (
	"context"
	"github.com/pkg/errors"

	"LingXi/app/images/cmd/rpc/internal/svc"
	"LingXi/app/images/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteImageLogic {
	return &DeleteImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteImageLogic) DeleteImage(in *pb.DeleteImageRequest) (*pb.DeleteImageResponse, error) {
	if err := l.svcCtx.ImageModel.Delete(l.ctx, nil, in.Id); err != nil {
		return nil, errors.Wrap(err, "failed to delete image")
	}

	return &pb.DeleteImageResponse{
		Id: in.GetId(),
	}, nil
}
