package posts

import (
	"LingXi/app/comment/cmd/rpc/commentcenter"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentsLogic {
	return &GetPostCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostCommentsLogic) GetPostComments(req *types.GetCommentsRequst) (resp *types.GetCommentsResponse, err error) {
	CommentsResp, err := l.svcCtx.CommentCenterRpc.GetComments(l.ctx, &commentcenter.GetCommentRequest{
		Type:  "post",
		OutId: req.PostId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.GetCommentsResponse{}
	comments := make([]types.PostComment, len(CommentsResp.Comment))
	for i, m := range CommentsResp.Comment {
		commentImageResp, _ := l.svcCtx.ImageCenter.GetAllImages(l.ctx, &imagecenter.GetAllImagesRequest{
			Type:  "comment",
			OutId: m.CommentId,
		})
		//fmt.Println(commentImageResp.Images)
		if commentImageResp != nil {
			images := make([]types.CommentImage, len(commentImageResp.Images))
			for j, n := range commentImageResp.Images {
				images[j] = types.CommentImage{
					CommentId: m.CommentId,
					ImageId:   n.Id,
					Url:       n.Url,
				}
			}
			comments[i] = types.PostComment{
				CommentId:       m.CommentId,
				PublisherId:     m.PublisherId,
				PubliserName:    m.PublisherName,
				PublisherAvatar: m.PublisherAvatar,
				CommentImage:    images,
				PostId:          m.OutId,
				Content:         m.Content,
				Time:            m.Time,
			}
		} else {
			comments[i] = types.PostComment{
				CommentId:       m.CommentId,
				PublisherId:     m.PublisherId,
				PubliserName:    m.PublisherName,
				PublisherAvatar: m.PublisherAvatar,
				//CommentImage:    images,
				PostId:  m.OutId,
				Content: m.Content,
				Time:    m.Time,
			}
		}

	}

	resp.PostComment = comments
	return resp, nil
}
