package transport

import (
	"context"
	"log"
	_ "mtsaudio/docs"
	"mtsaudio/internal/transport/handlers/httphandlers"
	"mtsaudio/internal/transport/handlers/middleware"
	"mtsaudio/internal/transport/handlers/websockethandlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	addr   string
	router *gin.Engine
}

func New(addr string) Server {
	return Server{
		addr:   addr,
		router: gin.Default(),
	}
}

func (s *Server) Run(ctx context.Context, hUc httphandlers.AuthUsecase, wsUc websockethandlers.WebsocketUsecase, tUc httphandlers.TrackUsecase) {
	//swagger route
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//auth routes
	httpHandler := httphandlers.New(hUc, tUc)
	wsHandler := websockethandlers.New(wsUc)

	//auth routes
	authRouts := s.router.Group("/", middleware.CheckAuthification())
	s.router.POST("/api/Account/SignIn", httpHandler.UserSignIn)
	s.router.POST("/api/Account/SignUp", httpHandler.UserSignUp)
	authRouts.POST("/api/Account/SignOut", httpHandler.UserSignOut)
	s.router.GET("/api/Account/RefreshTokens", httpHandler.RefreshTokens)
	authRouts.GET("/api/Account/Me", httpHandler.UserMyAccount)

	//track route
	authRouts.GET("/api/Tracks", httpHandler.GetTracks)

	//room handler
	authRouts.POST("/api/Room", httpHandler.CreateRoom)

	//websocket handlers
	s.router.GET("/ws/:roomId/chat", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")

		wsHandler.ChatConnect(roomId)(ctx.Writer, ctx.Request)
	})

	s.router.GET("/ws/:roomId/file", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")

		wsHandler.FileConnect(roomId)(ctx.Writer, ctx.Request)
	})

	s.router.GET("/ws/:roomId/track", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")

		wsHandler.TrackConnect(roomId)(ctx.Writer, ctx.Request)
	})

	srv := http.Server{
		Addr:    s.addr,
		Handler: s.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("failed to listen server")
		}
	}()

	//gracefull shutdown
	<-ctx.Done()
	log.Println("closing server gracefully...")
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := srv.Shutdown(ctxTimeout); err != nil {
		log.Println("failed to shutdown server gracefully")
	}
	log.Println("server closed gracefully")
}
