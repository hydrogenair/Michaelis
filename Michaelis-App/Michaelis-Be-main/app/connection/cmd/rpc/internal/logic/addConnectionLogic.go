package logic

import (
	"context"

	"LingXi/app/connection/cmd/rpc/internal/svc"
	"LingXi/app/connection/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConnectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddConnectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConnectionLogic {
	return &AddConnectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddConnectionLogic) AddConnection(in *pb.AddConnectionRequest) (*pb.AddConnectionResponse, error) {

	return &pb.AddConnectionResponse{}, nil
}
