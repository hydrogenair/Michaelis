package posts

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"
	"LingXi/app/post/cmd/rpc/postcenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryPostLogic {
	return &GetCategoryPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryPostLogic) GetCategoryPost(req *types.GetCategoryPostRequest) (resp *types.GetCategoryPostResponse, err error) {
	postResp, err := l.svcCtx.PostCenterRpc.GetSomeonePost(l.ctx, &postcenter.GetSomeonePostRequest{
		Category:    req.Category,
		PublisherId: -1,
	})
	fmt.Println(postResp)
	resp = &types.GetCategoryPostResponse{}
	var posts = make([]types.Post, len(postResp.Post))
	for i, m := range postResp.Post {
		images, _ := l.svcCtx.ImageCenter.GetAllImages(l.ctx, &imagecenter.GetAllImagesRequest{
			Type:  "post",
			OutId: m.PostId,
		})

		var Images = make([]types.PostImage, len(images.Images))
		for j, n := range images.Images {
			Images[j] = types.PostImage{
				PostId:  n.OutId,
				ImageId: n.Id,
				Url:     n.Url,
			}
		}

		userResp, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
			Id: m.PublisherId,
		})

		userAvatarResp, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
			Type:  "avatar",
			OutId: userResp.User.Id,
		})
		if len(images.Images) == 0 {
			posts[i] = types.Post{
				PostId:          m.PostId,
				PublisherId:     m.PublisherId,
				PublisherName:   userResp.User.Nickname,
				PublisherAvatar: userAvatarResp.Image.Url,
				Category:        m.Category,
				Content:         m.Content,
				ViewNum:         m.ViewNum,
				LikeNum:         m.LikeNum,
				CommentNum:      m.CommentNum,
				Time:            m.Time,
			}
		} else {
			posts[i] = types.Post{
				PostId:          m.PostId,
				PublisherId:     m.PublisherId,
				PublisherName:   userResp.User.Nickname,
				PublisherAvatar: userAvatarResp.Image.Url,
				PostImage:       Images,
				Category:        m.Category,
				Content:         m.Content,
				ViewNum:         m.ViewNum,
				LikeNum:         m.LikeNum,
				CommentNum:      m.CommentNum,
				Time:            m.Time,
			}
		}
	}

	resp.Post = posts
	return resp, nil
}
