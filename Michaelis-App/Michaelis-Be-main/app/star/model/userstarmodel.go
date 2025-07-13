package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserStarModel = (*customUserStarModel)(nil)

type (
	// UserStarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserStarModel.
	UserStarModel interface {
		userStarModel
		customUserStarLogicModel
	}

	customUserStarModel struct {
		*defaultUserStarModel
	}

	customUserStarLogicModel interface {
		FindAll(ctx context.Context, userId int64) ([]*UserStar, error)
	}
)

func (c customUserStarModel) FindAll(ctx context.Context, userId int64) ([]*UserStar, error) {
	var user_star []*UserStar
	var err error
	err = c.QueryNoCacheCtx(ctx, &user_star, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserStar{}).Where("`user_id` = ? ", userId).Find(&user_star).Error
	})
	if err != nil {
		return nil, err
	}
	return user_star, nil
}

// NewUserStarModel returns a model for the database table.
func NewUserStarModel(conn *gorm.DB, c cache.CacheConf) UserStarModel {
	return &customUserStarModel{
		defaultUserStarModel: newUserStarModel(conn, c),
	}
}

func (m *defaultUserStarModel) customCacheKeys(data *UserStar) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
