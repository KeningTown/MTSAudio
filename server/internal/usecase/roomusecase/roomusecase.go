package websocketusecase

type RoomRepository interface {
	// CreateRoom(roomId string) models.Room
	// FindRoomByUserId(userId int) models.Room
	// DeleteRoomByUserId(userId int)
}

type RoomUsecase struct {
	r RoomRepository
}

func New(r RoomRepository) RoomUsecase {
	return RoomUsecase{r}
}

// func (ru RoomUsecase) CreateRoom(userId int, roomId string) (entities.Room, error) {
// 	room := ru.r.FindRoomByUserId(userId)
// 	if room.Id != 0 {
// 		return entities.Room{}, fmt.Errorf("user already have room")
// 	}
// 	roomModel := ru.r.CreateRoom(roomId)
// 	return dto.RoomModelToEntite(roomModel), nil
// }
