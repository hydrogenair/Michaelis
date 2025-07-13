package posts

import (
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"
	"LingXi/app/post/cmd/rpc/postcenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailsLogic {
	return &GetPostDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostDetailsLogic) GetPostDetails(req *types.GetDetailRequest) (resp *types.GetDetailResponse, err error) {
	resp = &types.GetDetailResponse{}
	//fmt.Println("--------------")
	//fmt.Println(req.PostId)
	//fmt.Println("--------------")
	postResp, err := l.svcCtx.PostCenterRpc.GetPostDetail(l.ctx, &postcenter.GetPostDetailRequest{
		PostId: req.PostId,
	})

	if err != nil {
		return nil, err
	}
	userResp, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Id: postResp.Post.PublisherId,
	})

	imageResp, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
		Type:  "avatar",
		OutId: userResp.User.Id,
	})

	postImageResp, _ := l.svcCtx.ImageCenter.GetAllImages(l.ctx, &imagecenter.GetAllImagesRequest{
		Type:  "post",
		OutId: req.PostId,
	})

	images := make([]types.PostImage, len(postImageResp.Images))
	for i, m := range postImageResp.Images {
		images[i] = types.PostImage{
			PostId:  req.PostId,
			ImageId: m.OutId,
			Url:     m.Url,
		}
	}
	if len(postImageResp.Images) == 0 {
		resp.Post = types.Post{
			PostId:          postResp.Post.PostId,
			PublisherId:     postResp.Post.PublisherId,
			PublisherName:   userResp.User.Nickname,
			PublisherAvatar: imageResp.Image.Url,
			Category:        postResp.Post.Category,
			Content:         postResp.Post.Content,
			ViewNum:         postResp.Post.ViewNum,
			LikeNum:         postResp.Post.LikeNum,
			CommentNum:      postResp.Post.CommentNum,
			Time:            postResp.Post.Time,
		}
	} else {
		resp.Post = types.Post{
			PostId:          postResp.Post.PostId,
			PublisherId:     postResp.Post.PublisherId,
			PublisherName:   userResp.User.Nickname,
			PublisherAvatar: imageResp.Image.Url,
			Category:        postResp.Post.Category,
			Content:         postResp.Post.Content,
			ViewNum:         postResp.Post.ViewNum,
			LikeNum:         postResp.Post.LikeNum,
			CommentNum:      postResp.Post.CommentNum,
			PostImage:       images,
			Time:            postResp.Post.Time,
		}
	}

	return resp, nil
}
