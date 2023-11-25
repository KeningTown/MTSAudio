package models

type Room struct {
	Id     uint `gorm:"primaryKey"`
	UserId uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserId"`
}
