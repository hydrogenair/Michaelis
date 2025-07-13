package logic

import (
	"context"
	"github.com/pkg/errors"

	"LingXi/app/chat/cmd/rpc/internal/svc"
	"LingXi/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLastMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLastMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLastMessageLogic {
	return &GetLastMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLastMessageLogic) GetLastMessage(in *pb.LastMessageRequest) (*pb.LastMessageResponse, error) {
	message, err := l.svcCtx.MessageModel.FindLastMessage(l.ctx, in.GetSenderId(), in.GetReceiverId())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all messages")
	}

	return &pb.LastMessageResponse{
		Content: message.Content,
		Type:    message.Type,
		Time:    message.CreateTime.String(),
	}, nil
}
