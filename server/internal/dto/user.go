package dto

import (
	"mtsaudio/internal/database/models"
	"mtsaudio/internal/entities"
)

func UserEntitieToModels(user entities.User) models.User {
	return models.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}
}

func UserModelToEntitie(user models.User) entities.User {
	return entities.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}
}
