package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ImagesModel = (*customImagesModel)(nil)

type (
	// ImagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customImagesModel.
	ImagesModel interface {
		imagesModel
		customImagesLogicModel
	}

	customImagesModel struct {
		*defaultImagesModel
	}

	customImagesLogicModel interface {
		GetAllImages(ctx context.Context, Type string, OutId int64) ([]*Images, error)
		GetImage(ctx context.Context, Type string, OutId int64) (*Images, error)
	}
)

func (c customImagesModel) GetImage(ctx context.Context, Type string, OutId int64) (*Images, error) {
	var image *Images
	var err error
	err = c.QueryNoCacheCtx(ctx, &image, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Images{}).Where("`type` = ? and `out_id` = ?", Type, OutId).First(&image).Error
	})
	if err != nil {
		return nil, err
	}
	return image, nil
}

func (c customImagesModel) GetAllImages(ctx context.Context, Type string, OutId int64) ([]*Images, error) {
	var images []*Images
	var err error
	err = c.QueryNoCacheCtx(ctx, &images, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Images{}).Where("`type` = ? and `out_id` = ?", Type, OutId).Find(&images).Error
	})
	if err != nil {
		return nil, err
	}
	return images, nil
}

// NewImagesModel returns a model for the database table.
func NewImagesModel(conn *gorm.DB, c cache.CacheConf) ImagesModel {
	return &customImagesModel{
		defaultImagesModel: newImagesModel(conn, c),
	}
}

func (m *defaultImagesModel) customCacheKeys(data *Images) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
