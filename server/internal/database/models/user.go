package models

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null; unique" `
	Password string `gorm:"not null"`
}
