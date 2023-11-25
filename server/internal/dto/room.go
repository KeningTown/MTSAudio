package dto

import (
	"mtsaudio/internal/database/models"
	"mtsaudio/internal/entities"
)

func RoomEntiteToModel(room entities.Room) models.Room {
	return models.Room{
		Id:     room.Id,
		UserId: room.UserId,
	}
}

func RoomModelToEntite(room models.Room) entities.Room {
	return entities.Room{
		Id:     room.Id,
		UserId: room.UserId,
	}
}
