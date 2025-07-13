package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ StarsModel = (*customStarsModel)(nil)

type (
	// StarsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStarsModel.
	StarsModel interface {
		starsModel
		customStarsLogicModel
	}

	customStarsModel struct {
		*defaultStarsModel
	}

	customStarsLogicModel interface {
	}
)

// NewStarsModel returns a model for the database table.
func NewStarsModel(conn *gorm.DB, c cache.CacheConf) StarsModel {
	return &customStarsModel{
		defaultStarsModel: newStarsModel(conn, c),
	}
}

func (m *defaultStarsModel) customCacheKeys(data *Stars) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
