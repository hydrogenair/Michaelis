package logic

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"context"
	"fmt"

	"LingXi/app/connection/cmd/rpc/internal/svc"
	"LingXi/app/connection/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckConnectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckConnectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckConnectionLogic {
	return &CheckConnectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckConnectionLogic) CheckConnection(in *pb.CheckConnectionRequest) (*pb.CheckConnectionResponse, error) {
	user, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Id: in.UserId,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(user.User)
	if user.User.Identity == 0 {
		connection, err := l.svcCtx.ConnectionModel.GetUserConnection(l.ctx, in.GetUserId())
		if err != nil {
			return nil, err
		}
		if connection != nil {
			var ConnectionResp = make([]*pb.UserInfo, len(connection))
			for i, m := range connection {
				u, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
					Id: m.VolunteerId,
				})
				avatar, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
					Type:  "avatar",
					OutId: u.User.Id,
				})
				resp := &pb.UserInfo{
					UserId:   u.User.Id,
					UserName: u.User.Nickname,
					Avatar:   avatar.Image.Url,
				}
				ConnectionResp[i] = resp
			}
			return &pb.CheckConnectionResponse{
				Info: ConnectionResp,
			}, nil
		} else {
			return &pb.CheckConnectionResponse{}, nil
		}
	} else {
		connection, err := l.svcCtx.ConnectionModel.GetVolunteerConnection(l.ctx, in.GetUserId())
		if err != nil {
			return nil, err
		}
		if connection != nil {
			var ConnectionResp = make([]*pb.UserInfo, len(connection))
			for i, m := range connection {
				u, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
					Id: m.UserId,
				})
				avatar, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
					Type:  "avatar",
					OutId: u.User.Id,
				})
				resp := &pb.UserInfo{
					UserId:   u.User.Id,
					UserName: u.User.Nickname,
					Avatar:   avatar.Image.Url,
				}
				ConnectionResp[i] = resp
			}
			return &pb.CheckConnectionResponse{
				Info: ConnectionResp,
			}, nil
		} else {
			return &pb.CheckConnectionResponse{}, nil
		}
	}
}
