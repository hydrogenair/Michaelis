package chat

import (
	"LingXi/app/chat/cmd/api/internal/svc"
	"LingXi/app/chat/cmd/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.SocketReq) (resp *types.SocketResp, err error) {
	resp = &types.SocketResp{}
	return resp, nil
}
