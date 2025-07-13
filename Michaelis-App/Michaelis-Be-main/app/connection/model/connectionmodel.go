package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ConnectionModel = (*customConnectionModel)(nil)

type (
	// ConnectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConnectionModel.
	ConnectionModel interface {
		connectionModel
		customConnectionLogicModel
	}

	customConnectionModel struct {
		*defaultConnectionModel
	}

	customConnectionLogicModel interface {
		GetVolunteerConnection(ctx context.Context, ID int64) ([]*Connection, error)
		GetUserConnection(ctx context.Context, ID int64) ([]*Connection, error)
	}
)

func (c customConnectionModel) GetVolunteerConnection(ctx context.Context, ID int64) ([]*Connection, error) {
	var connections []*Connection
	var err error
	err = c.QueryNoCacheCtx(ctx, &connections, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Connection{}).Where("`volunteer_id`=?", ID).Find(&connections).Error
	})
	if err != nil {
		return nil, err
	}
	return connections, nil
}

func (c customConnectionModel) GetUserConnection(ctx context.Context, ID int64) ([]*Connection, error) {
	var connections []*Connection
	var err error
	err = c.QueryNoCacheCtx(ctx, &connections, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Connection{}).Where("`user_id`=?", ID).Find(&connections).Error
	})
	if err != nil {
		return nil, err
	}
	return connections, nil
}

// NewConnectionModel returns a model for the database table.
func NewConnectionModel(conn *gorm.DB, c cache.CacheConf) ConnectionModel {
	return &customConnectionModel{
		defaultConnectionModel: newConnectionModel(conn, c),
	}
}

func (m *defaultConnectionModel) customCacheKeys(data *Connection) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
