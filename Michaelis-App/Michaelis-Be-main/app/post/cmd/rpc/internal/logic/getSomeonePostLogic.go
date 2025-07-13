package logic

import (
	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"
	"LingXi/common/xerr"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSomeonePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSomeonePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSomeonePostLogic {
	return &GetSomeonePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSomeonePostLogic) GetSomeonePost(in *pb.GetSomeonePostRequest) (*pb.GetSomeonePostResponse, error) {

	posts, err := l.svcCtx.PostModel.FindAll(l.ctx, in.PublisherId, in.Category)

	if err != nil {
		return nil, xerr.NewErrMsg("Don't have any posts")
	}
	var Posts = make([]*pb.Post, len(posts))
	for i, m := range posts {
		//post := new(pb.Post)
		Posts[i] = &pb.Post{
			PostId:      m.Id,
			PublisherId: m.PublisherId,
			Content:     m.Content,
			Category:    m.Category,
			ViewNum:     m.ViewNum,
			LikeNum:     m.LikeNum,
			CommentNum:  m.CommentNum,
			Time:        m.CreateTime.String(),
		}
		//Posts = append(Posts, post)
	}

	return &pb.GetSomeonePostResponse{
		Post: Posts,
	}, nil
}
