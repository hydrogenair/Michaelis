package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CommentsModel = (*customCommentsModel)(nil)

type (
	// CommentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentsModel.
	CommentsModel interface {
		commentsModel
		customCommentsLogicModel
	}

	customCommentsModel struct {
		*defaultCommentsModel
	}

	customCommentsLogicModel interface {
		FindAll(ctx context.Context, Type string, OutId int64) ([]*Comments, error)
	}
)

func (c customCommentsModel) FindAll(ctx context.Context, Type string, OutId int64) ([]*Comments, error) {
	var comments []*Comments
	var err error
	err = c.QueryNoCacheCtx(ctx, &comments, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Comments{}).Where("`type` = ? and `out_id` = ?", Type, OutId).Find(&comments).Error
	})
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// NewCommentsModel returns a model for the database table.
func NewCommentsModel(conn *gorm.DB, c cache.CacheConf) CommentsModel {
	return &customCommentsModel{
		defaultCommentsModel: newCommentsModel(conn, c),
	}
}

func (m *defaultCommentsModel) customCacheKeys(data *Comments) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
