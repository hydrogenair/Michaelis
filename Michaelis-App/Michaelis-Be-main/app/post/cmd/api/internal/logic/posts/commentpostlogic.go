package posts

import (
	"LingXi/app/comment/cmd/rpc/commentcenter"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"LingXi/app/post/cmd/rpc/postcenter"
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPostLogic {
	return &CommentPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPostLogic) CommentPost(req *types.CommentPostRequest) (resp *types.CommentPostResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	postCommentResp, err := l.svcCtx.PostCenterRpc.CommentPost(l.ctx, &postcenter.CommentPostRequest{
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}
	commentResp, err := l.svcCtx.CommentCenterRpc.PostComment(l.ctx, &commentcenter.CommentRequest{
		PublisherId: userId,
		Content:     req.Content,
		Type:        "post",
		OutId:       req.PostId,
	})

	userResp, _ := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Id: userId,
	})
	userImageResp, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
		Type:  "avatar",
		OutId: userId,
	})
	resp = &types.CommentPostResponse{
		PostComment: types.PostComment{
			CommentId:       commentResp.Comment.CommentId,
			PublisherId:     commentResp.Comment.PublisherId,
			PubliserName:    userResp.User.Nickname,
			PublisherAvatar: userImageResp.Image.Url,
			PostId:          postCommentResp.PostId,
			Content:         commentResp.Comment.Content,
			Time:            commentResp.Comment.Time,
		},
	}
	return resp, nil
}
