package logic

import (
	"context"
	"github.com/pkg/errors"

	"LingXi/app/images/cmd/rpc/internal/svc"
	"LingXi/app/images/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDetailLogic) GetDetail(in *pb.GetDetailRequest) (*pb.GetDetailResponse, error) {
	detail, err := l.svcCtx.ImageModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get image detail")
	}

	return &pb.GetDetailResponse{
		Images: &pb.Image{
			Id:    detail.Id,
			Type:  detail.Type,
			OutId: detail.OutId,
			Url:   detail.Url,
		},
	}, nil
}
