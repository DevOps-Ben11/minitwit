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
	err := repo.db.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error

	if err != nil || user == nil {
		return nil, false
	}

	return user, true
}

func (repo UserRepository) GetUserById(user_id uint) (user *model.User, ok bool) {
	err := repo.db.Raw("SELECT * FROM users WHERE user_id = ?", user_id).Scan(&user).Error

	if err != nil || user == nil {
		return nil, false
	}

	return user, true
}

func (repo UserRepository) InsertUser(username string, email string, password string) error {
	return repo.db.Exec("INSERT INTO users (username, email, pw_hash) VALUES (?, ?, ?)",
		username, email, util.GeneratePasswordHash(password),
	).Error
}

func (repo UserRepository) GetUserTimeline(user_id uint) ([]model.RenderMessage, error) {
	var messages []model.RenderMessage

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
	followed := false
	repo.db.Raw("select 1 from follower where follower.who_id = ? and follower.whom_id = ?", who, whom).Scan(&followed)
	return followed
}

func (repo UserRepository) SetFollow(who uint, whom uint) error {
	err := repo.db.Exec("insert into follower (who_id, whom_id) values (?, ?)", who, whom).Error
	return err
}
func (repo UserRepository) SetUnfollow(who uint, whom uint) error {
	err := repo.db.Exec("delete from follower where who_id=? and whom_id=?", who, whom).Error
	return err
}

func (repo UserRepository) GetUsersFollowing(userId uint, limit int) ([]string, error) {
	var usernames []string
	err := repo.db.Raw(`
        SELECT users.username FROM users
                   INNER JOIN follower ON follower.whom_id=users.user_id
                   WHERE follower.who_id=?
                   LIMIT ?
        `, userId, limit).Scan(&usernames).Error
	return usernames, err
}
