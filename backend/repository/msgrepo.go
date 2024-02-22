package repository

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetPublicMessage() ([]model.RenderMessage, error)
}

type MessageRepository struct {
	db *gorm.DB
}
