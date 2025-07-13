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

type GetPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostLogic) GetPost(req *types.GetPostRequest) (resp *types.GetPostResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	fmt.Println(userId)
	resp = &types.GetPostResponse{}
	postResp, err := l.svcCtx.PostCenterRpc.GetSomeonePost(l.ctx, &postcenter.GetSomeonePostRequest{
		PublisherId: userId,
		Category:    "所有",
	})

	if err != nil {
		return nil, err
	}
	var posts = make([]types.Post, len(postResp.Post))
	userResp, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Id: userId,
	})

	userimageResp, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
		Type:  "avatar",
		OutId: userId,
	})

	for i, m := range postResp.Post {
		postImages, _ := l.svcCtx.ImageCenter.GetAllImages(l.ctx, &imagecenter.GetAllImagesRequest{
			Type:  "post",
			OutId: m.PostId,
		})
		images := make([]types.PostImage, len(postImages.Images))
		for j, n := range postImages.GetImages() {
			images[j] = types.PostImage{
				PostId:  n.OutId,
				ImageId: n.Id,
				Url:     n.Url,
			}
		}
		if len(images) == 0 {
			posts[i] = types.Post{
				PostId:          m.PostId,
				PublisherId:     m.PublisherId,
				PublisherAvatar: userimageResp.Image.Url,
				PublisherName:   userResp.User.Nickname,

				Category:   m.Category,
				Content:    m.Content,
				ViewNum:    m.ViewNum,
				LikeNum:    m.LikeNum,
				CommentNum: m.CommentNum,
				Time:       m.Time,
			}
		} else {
			posts[i] = types.Post{
				PostId:          m.PostId,
				PublisherId:     m.PublisherId,
				PublisherAvatar: userimageResp.Image.Url,
				PublisherName:   userResp.User.Nickname,
				PostImage:       images,
				Category:        m.Category,
				Content:         m.Content,
				ViewNum:         m.ViewNum,
				LikeNum:         m.LikeNum,
				CommentNum:      m.CommentNum,
				Time:            m.Time,
			}
		}

		//if err !=nil{
		//	continue
		//}
	}

	resp.Post = posts
	return resp, nil
}
