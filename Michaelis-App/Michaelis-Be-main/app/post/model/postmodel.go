package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		customPostLogicModel
	}

	customPostModel struct {
		*defaultPostModel
	}

	customPostLogicModel interface {
		FindAll(ctx context.Context, publisherId int64, category string) ([]*Post, error)
	}
)

func (c customPostModel) FindAll(ctx context.Context, publisherId int64, category string) ([]*Post, error) {
	var posts []*Post
	var err error
	//postIdKey := fmt.Sprintf("%s%v", cachePostIdPrefix, publisherId)
	if category == "所有" {
		err = c.QueryNoCacheCtx(ctx, &posts, func(conn *gorm.DB, v interface{}) error {
			return conn.Model(&Post{}).Where("`publisher_id` = ? and `delete_time` is null", publisherId).Find(&posts).Error
		})
		if err == nil {
			return posts, nil
		}
	}
	if publisherId == -1 {
		err = c.QueryNoCacheCtx(ctx, &posts, func(conn *gorm.DB, v interface{}) error {
			return conn.Model(&Post{}).Where("`category` = ? and `delete_time` is null ", category).Find(&posts).Error
		})
		if err == nil {
			return posts, nil
		}
	}

	err = c.QueryNoCacheCtx(ctx, &posts, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Post{}).Where("`publisher_id` = ? and `category` = ? and `delete_time` is null", publisherId, category).Find(&posts).Error
	})

	return posts, err
}

// NewPostModel returns a model for the database table.
func NewPostModel(conn *gorm.DB, c cache.CacheConf) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(conn, c),
	}
}

func (m *defaultPostModel) customCacheKeys(data *Post) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
