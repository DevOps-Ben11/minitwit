package repository

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gorm.io/gorm"
)

var (
	messagesAdd = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minitwit_messages",
		Help: "The total number of message.",
	})
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
	err := m.db.Model(&model.Message{}).Select("message.*, user.*").Joins("LEFT JOIN user ON user.user_id = message.author_id").Where("message.flagged = 0").Order("message.pub_date DESC").Limit(limit).Scan(&messages).Error

	return messages, err
}

func (m MessageRepository) GetUserMessages(userId uint, limit int) ([]model.RenderMessage, error) {
	var messages []model.RenderMessage
	err := m.db.Model(&model.Message{}).Select("message.*, user.*").Joins("LEFT JOIN user ON user.user_id = message.author_id").Where("user.user_id = ?", userId).Order("message.pub_date DESC").Limit(limit).Scan(&messages).Error
	return messages, err
}

func (m MessageRepository) AddMessage(user *model.User, text string) error {
	err := m.db.Create(&model.Message{
		Author_id: user.User_id,
		Text:      text,
		Flagged:   false,
	}).Error
	if err == nil {
		messagesAdd.Inc()
	}
	return err
}
