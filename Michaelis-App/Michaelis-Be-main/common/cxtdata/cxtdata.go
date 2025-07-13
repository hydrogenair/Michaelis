package ctxdata

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"
var CtxKeyJwtUserIdentity = "identity"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}

func GetIdentityFromCtx(ctx context.Context) int64 {
	var identity int64
	if jsonIdentity, ok := ctx.Value(CtxKeyJwtUserIdentity).(json.Number); ok {
		if int64Identity, err := jsonIdentity.Int64(); err == nil {
			identity = int64Identity
		} else {
			logx.WithContext(ctx).Errorf("GetIdentityFromCtx err : %+v", err)
		}
	}
	return identity
}
