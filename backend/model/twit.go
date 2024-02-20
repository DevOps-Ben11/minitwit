package model

type User struct {
	ID       uint
	Username string
	Email    string
	Password string
}

type Follower struct {
	Who_id  uint
	Whom_id uint
}

type Message struct {
	ID       uint
	UserID   uint
	Text     string
	Pub_date int64 `gorm:"autoCreateTime:milli"`
	Flagged  bool
}
