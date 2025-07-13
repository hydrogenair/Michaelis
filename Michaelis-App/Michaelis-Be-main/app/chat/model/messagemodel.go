package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		customMessageLogicModel
	}

	customMessageModel struct {
		*defaultMessageModel
	}

	customMessageLogicModel interface {
		FindAllMessage(ctx context.Context, SenderId int64, ReceiverId int64) ([]*Message, error)
		FindLastMessage(ctx context.Context, SenderId int64, ReceiverId int64) (*Message, error)
	}
)

func (c customMessageModel) FindLastMessage(ctx context.Context, SenderId int64, ReceiverId int64) (*Message, error) {
	var msg Message
	err := c.QueryNoCacheCtx(ctx, &msg, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Message{}).Where("(`sender_id` = ? and `receiver_id`= ?)  or (`sender_id` = ? and `receiver_id`= ? )", SenderId, ReceiverId, ReceiverId, SenderId).Order("create_time desc").First(&msg).Error
	})
	if err == nil {
		return &msg, nil
	}
	return nil, err
}

func (c customMessageModel) FindAllMessage(ctx context.Context, SenderId int64, ReceiverId int64) ([]*Message, error) {
	var msg []*Message
	err := c.QueryNoCacheCtx(ctx, &msg, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Message{}).Where("(`sender_id` = ? and `receiver_id`= ?)  or (`sender_id` = ? and `receiver_id`= ? )", SenderId, ReceiverId, ReceiverId, SenderId).Find(&msg).Error
	})
	if err == nil {
		return msg, nil
	}
	return nil, err
}

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn *gorm.DB, c cache.CacheConf) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn, c),
	}
}

func (m *defaultMessageModel) customCacheKeys(data *Message) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
