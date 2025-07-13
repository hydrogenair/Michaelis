package connection

import (
	"LingXi/app/chat/cmd/rpc/chatcenter"
	"LingXi/app/connection/cmd/rpc/connectioncenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/connection/cmd/api/internal/svc"
	"LingXi/app/connection/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckConnectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckConnectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckConnectionLogic {
	return &CheckConnectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckConnectionLogic) CheckConnection(req *types.GetConnectionRequest) (resp *types.GetConnectionResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	connectionResp, err := l.svcCtx.ConnectionCenter.CheckConnection(l.ctx, &connectioncenter.CheckConnectionRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	if connectionResp != nil {
		var info = make([]types.UserInfo, len(connectionResp.Info))
		for i, m := range connectionResp.Info {
			msg, _ := l.svcCtx.ChatCenter.GetLastMessage(l.ctx, &chatcenter.LastMessageRequest{
				SenderId:   userId,
				ReceiverId: m.GetUserId(),
			})
			if msg != nil {
				info[i] = types.UserInfo{
					UserId:   m.UserId,
					UserName: m.UserName,
					Avatar:   m.Avatar,
					LastMessage: types.LastMessage{
						Content: msg.Content,
						Type:    msg.Type,
						Time:    msg.Time,
					},
				}
			} else {
				info[i] = types.UserInfo{
					UserId:      m.UserId,
					UserName:    m.UserName,
					Avatar:      m.Avatar,
					LastMessage: types.LastMessage{},
				}
			}

		}
		return &types.GetConnectionResponse{
				UserInfo: info,
			},
			nil
	} else {
		return &types.GetConnectionResponse{}, nil
	}

}
