package svc

import (
	"LingXi/app/comment/cmd/rpc/commentcenter"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/post/cmd/api/internal/config"
	"LingXi/app/post/cmd/rpc/postcenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	PostCenterRpc    postcenter.PostCenter
	CommentCenterRpc commentcenter.CommentCenter
	ImageCenter      imagecenter.ImageCenter
	UserCenterRpc    usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		PostCenterRpc:    postcenter.NewPostCenter(zrpc.MustNewClient(c.PostcenterRpcConf)),
		CommentCenterRpc: commentcenter.NewCommentCenter(zrpc.MustNewClient(c.CommentcenterRpcConf)),
		ImageCenter:      imagecenter.NewImageCenter(zrpc.MustNewClient(c.ImagecenterRpcConf)),
		UserCenterRpc:    usercenter.NewUserCenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
