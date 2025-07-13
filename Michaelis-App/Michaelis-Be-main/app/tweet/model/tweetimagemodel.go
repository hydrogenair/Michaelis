package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TweetImageModel = (*customTweetImageModel)(nil)

type (
	// TweetImageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTweetImageModel.
	TweetImageModel interface {
		tweetImageModel
		customTweetImageLogicModel
	}

	customTweetImageModel struct {
		*defaultTweetImageModel
	}

	customTweetImageLogicModel interface {
	}
)

// NewTweetImageModel returns a model for the database table.
func NewTweetImageModel(conn *gorm.DB, c cache.CacheConf) TweetImageModel {
	return &customTweetImageModel{
		defaultTweetImageModel: newTweetImageModel(conn, c),
	}
}

func (m *defaultTweetImageModel) customCacheKeys(data *TweetImage) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
