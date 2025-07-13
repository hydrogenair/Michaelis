package logic

import (
	"LingXi/app/comment/cmd/rpc/internal/svc"
	"LingXi/app/comment/cmd/rpc/pb"
	"LingXi/app/images/cmd/rpc/imagecenter"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentsLogic) GetComments(in *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	comments, err := l.svcCtx.CommentModel.FindAll(l.ctx, in.Type, in.OutId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get comments")
	}
	//fmt.Println(comments[0])
	if comments == nil {
		return &pb.GetCommentResponse{}, nil
	}
	Comments := make([]*pb.Comment, len(comments))
	for i, m := range comments {
		//fmt.Println("----------")
		//fmt.Println(i)
		user, _ := l.svcCtx.UserModel.FindOne(l.ctx, m.PublisherId)
		//fmt.Println("user")
		//fmt.Println(user)
		userImageResp, _ := l.svcCtx.ImageCenter.GetOneImage(l.ctx, &imagecenter.GetOneImageRequest{
			Type:  "avatar",
			OutId: user.Id,
		})
		//fmt.Println("image")
		//fmt.Println(image)

		Comments[i] = &pb.Comment{
			CommentId:       m.Id,
			PublisherId:     m.PublisherId,
			PublisherName:   user.Nickname,
			PublisherAvatar: userImageResp.Image.Url,
			Type:            m.Type,
			Content:         m.Content,
			LikeNum:         m.LikeNum,
			ReplyNum:        m.ReplyNum,
			OutId:           m.OutId,
			Time:            m.CreateTime.String(),
		}
	}
	return &pb.GetCommentResponse{
		Comment: Comments,
	}, nil

}
