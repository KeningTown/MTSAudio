package models

type Token struct {
	Id     uint   `gorm:"primaryKey"`
	UserId uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserId"`
	Token  string `gorm:"not null"`
}
