package repository

import (
	"github.com/DevOps-Ben11/minitwit/backend/model"
	"github.com/DevOps-Ben11/minitwit/backend/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gorm.io/gorm"
)

var (
	usersRegistered = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minitwit_registrations",
		Help: "The total number of registrations.",
	})

	usersFollowed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minitwit_follows",
		Help: "The total number of follows.",
	})

	usersUnfollowed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minitwit_unfollows",
		Help: "The total number of unfollows.",
	})
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
  err := repo.db.Create(&model.User{Username: username, Email: email, Pw_hash: util.GeneratePasswordHash(password)}).Error
  if err == nil {
		usersRegistered.Inc()
	}
	return err
}

func (repo UserRepository) GetUserTimeline(user_id uint) ([]model.RenderMessage, error) {
	var messages []model.RenderMessage

	err := repo.db.Model(&model.Message{}).Model(&model.Message{}).Joins("LEFT JOIN users ON users.user_id = messages.author_id").Where("messages.flagged = ? AND (users.user_id = ? OR users.user_id IN (?))", false, user_id,
		repo.db.Model(&model.Follower{}).Where("who_id = ?", user_id).Select("whom_id")).Order("messages.pub_date DESC").Limit(util.PER_PAGE).Select("messages.*", "users.*").Find(&messages).Error

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
  if err == nil {
		usersUnfollowed.Inc()
	}
	return err
}

func (repo UserRepository) SetUnfollow(who uint, whom uint) error {
	err := repo.db.Delete(&model.Follower{}, "who_id=? and whom_id=?", who, whom).Error
	return err
}

func (repo UserRepository) GetUsersFollowing(userId uint, limit int) ([]string, error) {
	var usernames []string

	err := repo.db.Model(&model.User{}).Select("users.username").Joins("INNER JOIN followers ON followers.whom_id=users.user_id").Where("followers.who_id=?", userId).Limit(limit).Scan(&usernames).Error
	return usernames, err
}
