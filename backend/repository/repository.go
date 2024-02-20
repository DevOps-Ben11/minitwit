package repository

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"gorm.io/gorm"
)

type Repository interface {
	GetUser(username string) (*model.User, bool)
	InsertUser(username string, email string, password string) error
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

func (repo *repository) InsertUser(username string, email string, password string) error {
	return repo.db.Exec("INSERT INTO user (username, email, pw_hash) VALUES (?, ?, ?)",
		username, email, util.GeneratePasswordHash(password),
	).Error
}
