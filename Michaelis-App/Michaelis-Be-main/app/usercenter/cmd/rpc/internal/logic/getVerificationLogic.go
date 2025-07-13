package logic

import (
	"LingXi/common/helpers"
	"LingXi/common/mail"
	rds "LingXi/common/redis"
	"LingXi/common/verifycode"
	"LingXi/common/xerr"
	"context"
	"github.com/pkg/errors"

	"LingXi/app/usercenter/cmd/rpc/internal/svc"
	"LingXi/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVerificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVerificationLogic {
	return &GetVerificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVerificationLogic) GetVerification(in *pb.VerificationReq) (*pb.VerificationResp, error) {
	//u, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	//if u != nil {
	//	return nil, errors.Wrap(xerr.NewErrCode(xerr.EMAIL_REGISTERED_ERROR), "")
	//}
	code, err := l.getVerificationCode(in.Email)
	if err != nil {
		return nil, errors.Wrap(xerr.NewErrCode(100001), err.Error())
	}
	return &pb.VerificationResp{
		Verification: code,
	}, nil
}

func (l *GetVerificationLogic) getVerificationCode(email string) (string, error) {
	code := helpers.RandomNumber(l.svcCtx.Config.VerifyCode.CodeLength)

	var store = verifycode.RedisStore{
		RedisClient: &rds.RedisClient{
			Client:  l.svcCtx.RedisClient,
			Context: l.ctx,
		},
		KeyPrefix: l.svcCtx.Config.APP.Name + ":verifycode:",
	}
	if err := store.Set(email, code, l.svcCtx.Config.VerifyCode.ExpireTime); err != nil {
		return "", err
	}

	dia := mail.NewMailer(l.svcCtx.Config.Email.Host, l.svcCtx.Config.Email.Port,
		l.svcCtx.Config.Email.UserName, l.svcCtx.Config.Email.Password)
	if err := dia.Driver.Send(l.svcCtx.Config.Email.UserName,
		email, code); err != nil {
		return "", err
	}
	return code, nil
}
