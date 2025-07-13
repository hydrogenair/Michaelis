package connection

import (
	"context"

	"LingXi/app/connection/cmd/api/internal/svc"
	"LingXi/app/connection/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConnectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConnectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConnectionLogic {
	return &AddConnectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConnectionLogic) AddConnection(req *types.AddConnectionRequest) (resp *types.AddConnectionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
