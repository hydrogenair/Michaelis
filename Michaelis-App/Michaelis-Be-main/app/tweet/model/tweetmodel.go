package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TweetModel = (*customTweetModel)(nil)

type (
	// TweetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTweetModel.
	TweetModel interface {
		tweetModel
		customTweetLogicModel
	}

	customTweetModel struct {
		*defaultTweetModel
	}

	customTweetLogicModel interface {
	}
)

// NewTweetModel returns a model for the database table.
func NewTweetModel(conn *gorm.DB, c cache.CacheConf) TweetModel {
	return &customTweetModel{
		defaultTweetModel: newTweetModel(conn, c),
	}
}

func (m *defaultTweetModel) customCacheKeys(data *Tweet) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
