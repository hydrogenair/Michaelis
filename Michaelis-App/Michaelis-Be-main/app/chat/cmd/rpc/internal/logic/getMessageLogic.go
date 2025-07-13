package logic

import (
	"context"
	"github.com/pkg/errors"

	"LingXi/app/chat/cmd/rpc/internal/svc"
	"LingXi/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageLogic) GetMessage(in *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	//fmt.Println(in)
	message, err := l.svcCtx.MessageModel.FindAllMessage(l.ctx, in.GetSenderId(), in.GetReceiverId())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all messages")
	}
	var msg = make([]*pb.Message, len(message))
	for i, m := range message {
		msg[i] = &pb.Message{
			Id:         m.Id,
			SenderId:   m.SenderId,
			ReceiverId: m.ReceiverId,
			Type:       m.Type,
			Content:    m.Content,
			Time:       m.CreateTime.String(),
		}
	}
	return &pb.GetMessageResponse{
		Message: msg,
	}, nil
}
