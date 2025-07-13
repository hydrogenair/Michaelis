package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ DiaryModel = (*customDiaryModel)(nil)

type (
	// DiaryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDiaryModel.
	DiaryModel interface {
		diaryModel
		customDiaryLogicModel
	}

	customDiaryModel struct {
		*defaultDiaryModel
	}

	customDiaryLogicModel interface {
	}
)

// NewDiaryModel returns a model for the database table.
func NewDiaryModel(conn *gorm.DB, c cache.CacheConf) DiaryModel {
	return &customDiaryModel{
		defaultDiaryModel: newDiaryModel(conn, c),
	}
}

func (m *defaultDiaryModel) customCacheKeys(data *Diary) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
