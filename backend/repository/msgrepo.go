package repository

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetPublicMessages() ([]model.RenderMessage, error)
	AddMessage(user *model.User, text string) error
	GetUserMessages(user *model.User) ([]model.RenderMessage, error)
}

type MessageRepository struct {
	db *gorm.DB
}

func CreateMessageRepository(db *gorm.DB) IMessageRepository {
	return MessageRepository{db: db}
}

func (m MessageRepository) GetPublicMessages() ([]model.RenderMessage, error) {
	var messages []model.RenderMessage
	err := m.db.Raw(
		"select message.*, user.* from message, user where message.flagged = 0 and message.author_id = user.user_id order by message.pub_date desc limit ?",
		util.PER_PAGE).Scan(&messages).Error

	return messages, err
}

func (m MessageRepository) GetUserMessages(user *model.User) ([]model.RenderMessage, error) {
	var messages []model.RenderMessage
	err := m.db.Raw("select message.*, user.* from message, user where user.user_id = message.author_id and user.user_id = ? order by message.pub_date desc limit ?", user.User_id, util.PER_PAGE).Scan(&messages).Error
	return messages, err
}

func (m MessageRepository) AddMessage(user *model.User, text string) error {
	err := m.db.Create(&model.Message{
		Author_id: user.User_id,
		Text:      text,
		Flagged:   false,
	}).Error
	return err
}