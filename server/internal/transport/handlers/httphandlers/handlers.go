package httphandlers

import (
	"log"
	"mtsaudio/internal/entities"
	"mtsaudio/internal/transport/handlers/websockethandlers"
	"mtsaudio/internal/utils/httputils"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthUsecase interface {
	MyAccount(id uint) (entities.User, error)
	SignIn(user entities.User) (entities.User, string, string, error)
	SignUp(user entities.User) (entities.User, string, string, error)
	SignOut(token string)
	RefreshTokens(refreshToken string) (string, string, error)
}

type TrackUsecase interface {
	GetTracksName() []string
}

type HTTPHandler struct {
	hu AuthUsecase
	tu TrackUsecase
}

func New(hu AuthUsecase, tu TrackUsecase) HTTPHandler {
	return HTTPHandler{
		hu: hu,
		tu: tu,
	}
}

type ResponseUser struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

// @Summary Вход в аккаунт
// @Tags AccountController
// @Description Вход в аккаунт пользователя с использованием имени пользователя - username и паролем - password и получение jwt
// @Accept json
// @Produce  json
// @Param request body httphandlers.UserSignIn.userCreadentials true "User credentials"
// @Success 201 {object} httphandlers.UserSignIn.response
// @Failure 400 {object} httputils.ResponseError
// @Router /api/Account/SignIn [post]
func (hh HTTPHandler) UserSignIn(ctx *gin.Context) {
	type userCreadentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var userCred userCreadentials
	if err := ctx.BindJSON(&userCred); err != nil {
		httputils.NewResponseError(ctx, 400, err.Error())
		return
	}

	user := entities.User{Username: userCred.Username, Password: userCred.Password}
	user, accessToken, refreshToken, err := hh.hu.SignIn(user)
	if err != nil {
		httputils.NewResponseError(ctx, 400, err.Error())
		return
	}

	type response struct {
		User        ResponseUser `json:"user"`
		AccessToken string       `json:"access_token"`
	}

	log.Printf("user sign in: userId = %d username = %s", user.Id, user.Username)

	ctx.SetCookie("refresh_token", refreshToken, 2592000, "/", "localhost", false, true)
	ctx.JSON(201, response{
		User: ResponseUser{
			Id:       user.Id,
			Username: user.Username,
		},
		AccessToken: accessToken,
	})
}

// @Summary Обновление токенов
// @Tags AccountController
// @Description Обновление токенов и получение access_token в JSON и refresh_token в cookie
// @Produce  json
// @Success 201 {object} httphandlers.RefreshTokens.token
// @Failure 400 {object} httputils.ResponseError
// @Failure 401 {object} httputils.ResponseError
// @Router /api/Account/RefreshTokens [get]
func (hh HTTPHandler) RefreshTokens(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		httputils.NewResponseError(ctx, 400, "failed to get refresh token from cookie")
	}

	accessToken, refreshToken, err := hh.hu.RefreshTokens(refreshToken)
	if err != nil {
		httputils.NewResponseError(ctx, 400, err.Error())
		return
	}

	type token struct {
		AccessToken string `json:"access_token"`
	}

	ctx.SetCookie("refresh_token", refreshToken, 2592000, "/", "localhost", false, true)
	ctx.JSON(201, token{AccessToken: accessToken})
}

// @Summary Регистрация
// @Tags AccountController
// @Description Регистрация пользовате и получение jwt
// @Accept json
// @Produce  json
// @Param request body httphandlers.UserSignUp.userData true "User data"
// @Success 201 {object} httphandlers.UserSignIn.response
// @Failure 400 {object} httputils.ResponseError
// @Router /api/Account/SignUp [post]
func (hh HTTPHandler) UserSignUp(ctx *gin.Context) {
	type userData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var usData userData
	if err := ctx.BindJSON(&usData); err != nil {
		httputils.NewResponseError(ctx, 400, err.Error())
		return
	}
	user := entities.User{
		Username: usData.Username,
		Password: usData.Password,
	}

	user, accessToken, refreshToken, err := hh.hu.SignUp(user)
	if err != nil {
		httputils.NewResponseError(ctx, 400, err.Error())
		return
	}

	type response struct {
		// required: true
		User        ResponseUser `json:"user"`
		AccessToken string       `json:"access_token"`
	}

	log.Printf("sign up new user: userId = %d username = %s", user.Id, user.Username)

	ctx.SetCookie("refresh_token", refreshToken, 2592000, "/", "localhost", false, true)
	ctx.JSON(201, response{
		User: ResponseUser{
			Id:       user.Id,
			Username: user.Username,
		},
		AccessToken: accessToken,
	})
}

// @Summary Выход из аккаунта
// @Tags AccountController
// @Description Внесение текущего используемого токена доступа в черный список токенов
// @Security ApiKeyAuth
// @Success 200
// @Failure 400 {object} httputils.ResponseError
// @Failure 401 {object} httputils.ResponseError
// @Router /api/Account/SignOut [post]
func (hh HTTPHandler) UserSignOut(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	hh.hu.SignOut(token)
	ctx.SetCookie("refresh_token", "", 1, "/", "localhost", false, true)
	ctx.Status(200)
}

// @Summary Просмотр данных текущего аккаунта
// @Tags AccountController
// @Description Просмотр информации о текущем авторизованном аккаунте
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} entities.User
// @Failure 400 {object} httputils.ResponseError
// @Failure 401 {object} httputils.ResponseError
// @Router /api/Account/Me [get]
func (hh HTTPHandler) UserMyAccount(ctx *gin.Context) {
	id := ctx.GetUint("id")
	user, err := hh.hu.MyAccount(id)
	if err != nil {
		httputils.NewResponseError(ctx, 400, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Создание новой комнаты
// @Tags RoomController
// @Description Создание новой комнаты и получение ее uuid
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} httphandlers.CreateRoom.response
// @Failure 400 {object} httputils.ResponseError
// @Failure 401 {object} httputils.ResponseError
// @Router /api/Room [post]
func (hh HTTPHandler) CreateRoom(ctx *gin.Context) {
	id := ctx.GetUint("id")

	roomId := uuid.New().String()

	chatMu := &sync.RWMutex{}
	websockethandlers.ChatRooms[roomId] = websockethandlers.Room{
		OwnerId: id,
		Clients: make(map[websockethandlers.Client]struct{}),
		Mu:      chatMu,
	}

	fileMu := &sync.RWMutex{}
	websockethandlers.FileRooms[roomId] = websockethandlers.Room{
		OwnerId: id,
		Clients: make(map[websockethandlers.Client]struct{}),
		Mu:      fileMu,
	}

	stopTrackMu := &sync.RWMutex{}
	websockethandlers.TrackRooms[roomId] = websockethandlers.Room{
		OwnerId: id,
		Clients: make(map[websockethandlers.Client]struct{}),
		Mu:      stopTrackMu,
	}

	type response struct {
		RoomId string `json:"roomId"`
	}

	log.Printf("created new room: roomId = %s ownerId = %d", roomId, id)

	ctx.JSON(http.StatusOK, response{RoomId: roomId})
}

// @Summary Информация о треках
// @Tags TrackController
// @Description Создание новой комнаты и получение ее uuid
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} string
// @Failure 401 {object} httputils.ResponseError
// @Router /api/Tracks [get]
func (hh HTTPHandler) GetTracks(ctx *gin.Context) {
	tracks := hh.tu.GetTracksName()

	ctx.JSON(200, tracks)
}
