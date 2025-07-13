package logic

import (
	"LingXi/app/post/cmd/rpc/internal/svc"
	"LingXi/app/post/cmd/rpc/pb"
	"LingXi/common/xerr"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostDetailLogic) GetPostDetail(in *pb.GetPostDetailRequest) (*pb.GetPostDetailResponse, error) {
	detail, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查找帖子详细失败")
	}

	return &pb.GetPostDetailResponse{
		Post: &pb.Post{
			PostId:      detail.Id,
			PublisherId: detail.PublisherId,
			Content:     detail.Content,
			Category:    detail.Category,
			ViewNum:     detail.ViewNum,
			LikeNum:     detail.LikeNum,
			CommentNum:  detail.CommentNum,
			Time:        detail.CreateTime.String(),
		},
	}, nil
}
