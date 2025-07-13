package logic

import (
	"LingXi/app/chat/model"
	"context"
	"github.com/pkg/errors"
	"time"

	"LingXi/app/chat/cmd/rpc/internal/svc"
	"LingXi/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveMessageRequestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveMessageRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveMessageRequestLogic {
	return &SaveMessageRequestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveMessageRequestLogic) SaveMessageRequest(in *pb.SaveMessageRequest) (*pb.SaveMessageResponse, error) {
	var msg = model.Message{
		SenderId:   in.SenderId,
		ReceiverId: in.ReceiverId,
		Type:       in.Type,
		Content:    in.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		//DeleteTime: sql.NullTime{},
	}

	if err := l.svcCtx.MessageModel.Insert(l.ctx, nil, &msg); err != nil {
		return nil, errors.Wrap(err, "failed to save message")
	}

	return &pb.SaveMessageResponse{
		Message: &pb.Message{
			Id:         msg.Id,
			SenderId:   msg.SenderId,
			ReceiverId: msg.ReceiverId,
			Type:       msg.Type,
			Content:    msg.Content,
			Time:       msg.CreateTime.String(),
		},
	}, nil
}
