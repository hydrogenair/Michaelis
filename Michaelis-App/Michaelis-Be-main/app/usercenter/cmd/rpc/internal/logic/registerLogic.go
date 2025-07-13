package logic

import (
	"LingXi/app/usercenter/cmd/rpc/usercenter"
	"LingXi/app/usercenter/model"
	rds "LingXi/common/redis"
	"LingXi/common/tool"
	"LingXi/common/verifycode"
	"LingXi/common/xerr"
	"context"
	"time"

	"LingXi/app/usercenter/cmd/rpc/internal/svc"
	"LingXi/app/usercenter/cmd/rpc/pb"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")
var ErrVerificationError = xerr.NewErrMsg("verification wrong")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	if _, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email); err == nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Email, err)
	}

	if !l.Verify(in.Verification, in.Email) {
		return nil, errors.Wrapf(ErrVerificationError, "The code for verification does not match email:%s", in.Email)
	}
	user := model.User{
		Nickname:   tool.Krand(8, tool.KC_RAND_KIND_ALL),
		Email:      in.Email,
		Gender:     "无",
		Password:   tool.Md5ByString(in.Password),
		Identity:   0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err := l.svcCtx.UserModel.Insert(l.ctx, nil, &user); err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
	}
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userEmail : %s", in.Email)
	}
	return &pb.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *RegisterLogic) Verify(code, email string) bool {
	var store = verifycode.RedisStore{
		RedisClient: &rds.RedisClient{
			Client:  l.svcCtx.RedisClient,
			Context: l.ctx,
		},
		KeyPrefix: l.svcCtx.Config.APP.Name + ":verifycode:",
	}
	return store.Verify(email, code)
}
