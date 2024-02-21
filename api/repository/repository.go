package repository

import (
	"github.com/DevOps-Ben11/minitwit/api/model"
	"github.com/DevOps-Ben11/minitwit/api/util"
	"gorm.io/gorm"
)

type Repository interface {
	GetUser(username string) (*model.User, bool)
	GetUserById(user_id *uint) (*model.User, bool)
	InsertUser(username string, email string, password string) error
	GetUserTimeline(user_id uint) ([]model.RenderMessage, bool)
}

type repository struct {
	db *gorm.DB
}

func CreateRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) GetUser(username string) (*model.User, bool) {
	var user *model.User

	err := repo.db.Raw("SELECT * FROM user WHERE username = ?", username).Scan(&user).Error

	if err != nil || user == nil {
		return nil, false
	}

	return user, true
}

func (repo *repository) GetUserById(user_id *uint) (*model.User, bool) {
	var user *model.User

	err := repo.db.Raw("SELECT * FROM user WHERE user_id = ?", user_id).Scan(&user).Error

	if err != nil || user == nil {
		return nil, false
	}

	return user, true
}

func (repo *repository) InsertUser(username string, email string, password string) error {
	return repo.db.Exec("INSERT INTO user (username, email, pw_hash) VALUES (?, ?, ?)",
		username, email, util.GeneratePasswordHash(password),
	).Error
}

func (repo *repository) GetUserTimeline(user_id uint) ([]model.RenderMessage, bool) {
	var messages []model.RenderMessage

	err := repo.db.Raw(
		`SELECT message.*, user.* FROM message, user
			WHERE message.flagged = 0 AND message.author_id = user.user_id
				AND (user.user_id = ? OR user.user_id IN
					(SELECT whom_id FROM follower WHERE who_id = ?))
			ORDER BY message.pub_date DESC LIMIT ?
		`, user_id, user_id, util.PER_PAGE).Scan(&messages).Error

	if err != nil {
		return nil, false
	}

	return messages, true
}
