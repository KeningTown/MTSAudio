package authusecase

import (
	"fmt"
	"mtsaudio/internal/database/models"
	"mtsaudio/internal/dto"
	"mtsaudio/internal/entities"
	"mtsaudio/internal/tokens"
	"time"
)

type AuthRepository interface {
	FindUserByUsername(username string) models.User
	FindUserById(id uint) models.User
	CreateUser(user models.User) models.User
}

type AuthUsecase struct {
	r AuthRepository
}

func New(r AuthRepository) AuthUsecase {
	return AuthUsecase{r: r}
}

func (au AuthUsecase) MyAccount(id uint) (entities.User, error) {
	user := au.r.FindUserById(id)
	if user.Id == 0 {
		return entities.User{}, fmt.Errorf("user is not exist")
	}
	return dto.UserModelToEntitie(user), nil
}

func (au AuthUsecase) SignIn(user entities.User) (entities.User, string, string, error) {
	userModel := au.r.FindUserByUsername(user.Username)
	if userModel.Id == 0 {
		return entities.User{}, "", "", fmt.Errorf("username is not exist")
	}

	if user.Password != userModel.Password {
		return entities.User{}, "", "", fmt.Errorf("invalid password")
	}

	userEntite := dto.UserModelToEntitie(userModel)
	accessToken, refreshToken, err := tokens.GenerateNewJwts(userEntite)
	if err != nil {
		return entities.User{}, "", "", err
	}

	return dto.UserModelToEntitie(userModel), accessToken, refreshToken, nil
}

func (au AuthUsecase) SignUp(user entities.User) (entities.User, string, string, error) {
	candidate := au.r.FindUserByUsername(user.Username)
	if candidate.Id != 0 {
		return entities.User{}, "", "", fmt.Errorf("user is already exist")
	}

	userModel := dto.UserEntitieToModels(user)
	userModel = au.r.CreateUser(userModel)
	userEntite := dto.UserModelToEntitie(userModel)

	accesToken, refreshToken, err := tokens.GenerateNewJwts(userEntite)
	if err != nil {
		return entities.User{}, "", "", err
	}
	return userEntite, accesToken, refreshToken, nil
}

func (au AuthUsecase) SignOut(token string) {
	tokens.RemoveToken(token)
}

func (au AuthUsecase) RefreshTokens(accessToken, refreshToken string) (string, string, error) {
	refreshTokenEntitie, err := tokens.ParseToken(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if int(time.Now().Unix())-refreshTokenEntitie.ExpIn > refreshTokenEntitie.CreatedAt {
		return "", "", fmt.Errorf("refresh token is expired")
	}

	userModel := au.r.FindUserById(refreshTokenEntitie.Id)
	userEntite := dto.UserModelToEntitie(userModel)

	accesToken, refreshToken, err := tokens.GenerateNewJwts(userEntite)
	if err != nil {
		return "", "", err
	}

	return accesToken, refreshToken, nil
}
