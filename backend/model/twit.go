package model

type User struct {
	User_id  uint   `gorm:"primaryKey"`
	Username string `gorm:"index"`
	Email    string
	Pw_hash  string
}

type Message struct {
	Message_id uint `gorm:"primaryKey"`
	Author_id  uint `gorm:"index;index:composite"`
	Text       string
	Pub_date   int64 `gorm:"autoCreateTime;index;index:composite,sort:desc"`
	Flagged    bool  `gorm:"index:composite"`
}

type Follower struct {
	Who_id  uint `gorm:"primaryKey;autoIncrement:false;index"`
	Whom_id uint `gorm:"primaryKey;autoIncrement:false"`
}
