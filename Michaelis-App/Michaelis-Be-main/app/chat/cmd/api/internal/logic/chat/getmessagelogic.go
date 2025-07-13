package chat

import (
	"LingXi/app/chat/cmd/rpc/chatcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/chat/cmd/api/internal/svc"
	"LingXi/app/chat/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageLogic) GetMessage(req *types.GetMessageRequest) (resp *types.GetMessageResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	msgResp, err := l.svcCtx.ChatCenterRpc.GetMessage(l.ctx, &chatcenter.GetMessageRequest{
		SenderId:   userId,
		ReceiverId: req.ReceiverId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetMessageResponse{}
	Msg := make([]types.Message, len(msgResp.Message))
	for i, m := range msgResp.Message {
		Msg[i] = types.Message{
			Id:         m.Id,
			ReceiverId: m.ReceiverId,
			SenderId:   m.SenderId,
			Type:       m.Type,
			Content:    m.Content,
			Time:       m.Time,
		}
	}
	resp.Message = Msg
	return resp, nil
}
