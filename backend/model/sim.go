package model

type KeyVal struct {
	Key   string `gorm:"primaryKey"`
	Value string
}
