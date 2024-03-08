package repository

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUser(username string) (*model.User, bool)
	GetUserById(userId uint) (*model.User, bool)
	InsertUser(username string, email string, password string) error
	GetUserTimeline(userId uint) ([]model.RenderMessage, error)
	GetIsFollowing(who uint, whom uint) bool
	SetFollow(who uint, whom uint) error
	SetUnfollow(who uint, whom uint) error
	GetUsersFollowing(userId uint, limit int) ([]string, error)
}

type UserRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) IUserRepository {
	return UserRepository{db: db}
}

func (repo UserRepository) GetUser(username string) (user *model.User, ok bool) {
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil || user == nil {
		return nil, false
	}

	return user, true
}

func (repo UserRepository) GetUserById(user_id uint) (user *model.User, ok bool) {
	err := repo.db.Where("user_id = ?", user_id).First(&user).Error

	if err != nil || user == nil {
		return nil, false
	}

	return user, true
}

func (repo UserRepository) InsertUser(username string, email string, password string) error {
	return repo.db.Create(&model.User{Username: username, Email: email, Pw_hash: util.GeneratePasswordHash(password)}).Error
}

func (repo UserRepository) GetUserTimeline(user_id uint) ([]model.RenderMessage, error) {
	var messages []model.RenderMessage

	// We chose to not use the GORM query, since its less readable than the raw SQlite Query.
	// err := repo.db.Where("message.flagged=0 AND message.author_id = user.user_id AND (user.user_id = ? OR user.user_id IN (?)", user_id,
	// repo.db.Table("follower").Where("who_id = ?", user_id).Select("whom_id")).Order("message.pub_date DESC").Limit(util.PER_PAGE).Select("message.*", "user.*").Find(&messages).Error

	err := repo.db.Raw(
		`SELECT message.*, users.* FROM message, users
			WHERE message.flagged = 0 AND message.author_id = users.user_id
				AND (users.user_id = ? OR users.user_id IN
					(SELECT whom_id FROM follower WHERE who_id = ?))
			ORDER BY message.pub_date DESC LIMIT ?
		`, user_id, user_id, util.PER_PAGE).Scan(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}
func (repo UserRepository) GetIsFollowing(who uint, whom uint) bool {
	var f model.Follower
	// If first cannot find a value, this query will throw an error.
	err := repo.db.Where("who_id = ? and whom_id = ?", who, whom).First(&f).Error
	// The error is used to see if there is a following between who and whom. if there is no error returns true, otherwise returns false
	return err == nil
}

func (repo UserRepository) SetFollow(who uint, whom uint) error {
	err := repo.db.Create(&model.Follower{Who_id: who, Whom_id: whom}).Error
	return err
}
func (repo UserRepository) SetUnfollow(who uint, whom uint) error {
	err := repo.db.Delete(&model.Follower{}, "who_id=? and whom_id=?", who, whom).Error
	return err
}

func (repo UserRepository) GetUsersFollowing(userId uint, limit int) ([]string, error) {
	var usernames []string

	err := repo.db.Model(&model.User{}).Select("user.username").Joins("INNER JOIN follower ON follower.whom_id=user.user_id").Where("follower.who_id=?", userId).Limit(limit).Scan(&usernames).Error
	return usernames, err
}
