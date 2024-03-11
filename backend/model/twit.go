package model

type User struct {
	User_id  uint `gorm:"primaryKey"`
	Username string
	Email    string
	Pw_hash  string
}

type Message struct {
	Message_id uint `gorm:"primaryKey"`
	Author_id  uint
	Text       string
	Pub_date   int64 `gorm:"autoCreateTime"`
	Flagged    bool
}

type Follower struct {
	Who_id  uint `gorm:"primaryKey;autoIncrement:false"`
	Whom_id uint `gorm:"primaryKey;autoIncrement:false"`
}
