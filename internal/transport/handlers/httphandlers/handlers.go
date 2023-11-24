package httphandlers

import (
	"mtsaudio/internal/entities"
	"mtsaudio/internal/utils/httputils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthUsecase interface {
	MyAccount(id uint) (entities.User, error)
	SignIn(user entities.User) (entities.User, string, string, error)
	SignUp(user entities.User) (entities.User, string, string, error)
	SignOut(token string)
	RefreshTokens(refreshToken string) (string, string, error)
}

type HTTPHandler struct {
	hu AuthUsecase
}

func New(hu AuthUsecase) HTTPHandler {
	return HTTPHandler{hu: hu}
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
		User        entities.User `json:"user"`
		AccessToken string        `json:"access_token"`
	}

	ctx.SetCookie("refresh_token", refreshToken, 2592000, "/", "localhost", false, true)
	ctx.JSON(201, response{
		User:        user,
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
// @Success 201 {object} entities.User
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
		User        entities.User `json:"user"`
		AccessToken string        `json:"access_token"`
	}

	ctx.SetCookie("refresh_token", refreshToken, 2592000, "/", "localhost", false, true)
	ctx.JSON(201, response{
		User:        user,
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
