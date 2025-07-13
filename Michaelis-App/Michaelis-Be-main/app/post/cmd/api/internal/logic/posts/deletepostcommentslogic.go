package posts

import (
	"LingXi/app/comment/cmd/rpc/commentcenter"
	ctxdata "LingXi/common/cxtdata"
	"context"

	"LingXi/app/post/cmd/api/internal/svc"
	"LingXi/app/post/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostCommentsLogic {
	return &DeletePostCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostCommentsLogic) DeletePostComments(req *types.DeleteCommentRequest) (resp *types.DeleteCommentResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	deleteResp, err := l.svcCtx.CommentCenterRpc.DeleteComment(l.ctx, &commentcenter.DeleteCommentRequest{
		PublisherId: userId,
		CommentId:   req.CommentId,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.DeleteCommentResponse{
		CommentId: deleteResp.CommentId,
	}
	return resp, nil
}
