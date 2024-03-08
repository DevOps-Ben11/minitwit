package repository

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetPublicMessages(limit int) ([]model.RenderMessage, error)
	AddMessage(user *model.User, text string) error
	GetUserMessages(userId uint, limit int) ([]model.RenderMessage, error)
}

type MessageRepository struct {
	db *gorm.DB
}

func CreateMessageRepository(db *gorm.DB) IMessageRepository {
	return MessageRepository{db: db}
}

func (m MessageRepository) GetPublicMessages(limit int) ([]model.RenderMessage, error) {
	var messages []model.RenderMessage
	err := m.db.Model(&model.Message{}).Select("*").Joins("LEFT JOIN users ON users.user_id = messages.author_id").Where("messages.flagged = ?", false).Order("messages.pub_date DESC").Limit(limit).Scan(&messages).Error

	return messages, err
}

func (m MessageRepository) GetUserMessages(userId uint, limit int) ([]model.RenderMessage, error) {
	var messages []model.RenderMessage
	err := m.db.Model(&model.Message{}).Select("messages.*, users.*").Joins("LEFT JOIN users ON users.user_id = messages.author_id").Where("users.user_id = ?", userId).Order("messages.pub_date DESC").Limit(limit).Scan(&messages).Error
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
