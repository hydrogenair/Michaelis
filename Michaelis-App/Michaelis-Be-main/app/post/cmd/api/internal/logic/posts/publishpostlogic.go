package posts

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/post/cmd/rpc/postcenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	ctxdata "LingXi/common/cxtdata"
	"context"
	"fmt"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishPostLogic {
	return &PublishPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishPostLogic) PublishPost(req *types.PostRequest) (resp *types.PostResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(req)
	postResp, err := l.svcCtx.PostCenterRpc.SendPost(l.ctx, &postcenter.PostRequest{
		PublisherId: userId,
		Content:     req.Content,
		Category:    req.Category,
	})
	if err != nil {
		return nil, err
	}
	userResp, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Id: userId,
	})

	imageResp, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
		Type:  "avatar",
		OutId: userId,
	})
	//if err != nil {
	//	return nil, err
	//}
	
	return &types.PostResponse{
		Post: types.Post{
			PostId:          postResp.Post.PostId,
			PublisherId:     postResp.Post.PublisherId,
			PublisherName:   userResp.User.Nickname,
			PublisherAvatar: imageResp.Image.Url,
			Category:        postResp.Post.Category,
			Content:         postResp.Post.Content,
			ViewNum:         postResp.Post.ViewNum,
			LikeNum:         postResp.Post.LikeNum,
			CommentNum:      postResp.Post.CommentNum,
		},
	}, nil
}
