package posts

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"
	"LingXi/app/post/cmd/rpc/postcenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonalCategoryPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalCategoryPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalCategoryPostLogic {
	return &GetPersonalCategoryPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalCategoryPostLogic) GetPersonalCategoryPost(req *types.PersonalPostRequest) (resp *types.PersonalPostResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//fmt.Println(req)
	postResp, err := l.svcCtx.PostCenterRpc.GetSomeonePost(l.ctx, &postcenter.GetSomeonePostRequest{
		Category:    req.Category,
		PublisherId: userId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.PersonalPostResponse{}
	var posts = make([]types.Post, len(postResp.Post))
	for i, m := range postResp.Post {
		userResp, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
			Id: m.PublisherId,
		})
		avatarResp, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
			OutId: userResp.User.Id,
			Type:  "avatar",
		})
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
		if len(postImages.Images) == 0 {
			posts[i] = types.Post{
				PostId:          m.PostId,
				PublisherId:     m.PublisherId,
				PublisherName:   userResp.User.GetNickname(),
				PublisherAvatar: avatarResp.Image.Url,
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
				PublisherName:   userResp.User.GetNickname(),
				PublisherAvatar: avatarResp.Image.Url,
				PostImage:       images,
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
