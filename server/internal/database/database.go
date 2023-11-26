package database

import (
	"fmt"
	"log"
	"mtsaudio/internal/config"
	"mtsaudio/internal/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Connect(cfg *config.Config) (Database, error) {
	op := "database.Connect()"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s ",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		return Database{}, fmt.Errorf("%s: failed to connect to postgres: %w", op, err)
	}

	if err := db.AutoMigrate(&models.Token{}, models.User{}); err != nil {
		return Database{}, fmt.Errorf("%s: failed to migrate database: %w", op, err)
	}

	log.Println("succesfully migrate database")
	return Database{db: db}, nil
}

// auth repository
func (db Database) FindUserByUsername(username string) models.User {
	var user models.User
	db.db.Find(&user, "username=?", username)
	return user
}

func (db Database) FindUserById(id uint) models.User {
	var user models.User
	db.db.Find(&user, "id=?", id)
	return user
}

func (db Database) CreateUser(user models.User) models.User {
	db.db.Create(&user)
	return user
}

func (db Database) CreateToken(token string, userId int) models.Token {
	tokenModel := models.Token{
		UserId: uint(userId),
		Token:  token,
	}
	db.db.Create(&tokenModel)
	return tokenModel
}

func (db Database) FindTokenById(tokenId int) models.Token {
	var token models.Token
	db.db.Find(&token, "id = ?", tokenId)
	return token
}

func (db Database) DeleteToken(token string) {
	var tokenModel models.Token
	db.db.Delete(&tokenModel, "token = ?", token)
}
