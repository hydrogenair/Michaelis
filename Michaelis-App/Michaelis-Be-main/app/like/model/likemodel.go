package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LikeModel = (*customLikeModel)(nil)

type (
	// LikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikeModel.
	LikeModel interface {
		likeModel
		customLikeLogicModel
	}

	customLikeModel struct {
		*defaultLikeModel
	}

	customLikeLogicModel interface {
		FindByTpOutId(ctx context.Context, Type string, OutId int64) ([]*Like, error)
	}
)

func (c customLikeModel) FindByTpOutId(ctx context.Context, Type string, OutId int64) ([]*Like, error) {
	var like []*Like
	var err error
	err = c.QueryNoCacheCtx(ctx, &like, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Like{}).Where("`type`=? and `out_id`=?", Type, OutId).Find(&like).Error
	})
	if err != nil {
		return nil, err
	}
	return like, nil
}

// NewLikeModel returns a model for the database table.
func NewLikeModel(conn *gorm.DB, c cache.CacheConf) LikeModel {
	return &customLikeModel{
		defaultLikeModel: newLikeModel(conn, c),
	}
}

func (m *defaultLikeModel) customCacheKeys(data *Like) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
