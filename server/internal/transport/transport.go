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

<<<<<<< HEAD
=======
	"github.com/gin-contrib/cors"
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
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

<<<<<<< HEAD
func (s *Server) Run(ctx context.Context, huc httphandlers.AuthUsecase, wsUc websockethandlers.WebsocketUsecase) {
=======
func (s *Server) Run(ctx context.Context, hUc httphandlers.AuthUsecase, wsUc websockethandlers.WebsocketUsecase, tUc httphandlers.TrackUsecase) {
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
	//swagger route
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//auth routes
<<<<<<< HEAD
	httpHandler := httphandlers.New(huc)
	wsHandler := websockethandlers.New(wsUc)

	//user auth routes
=======
	httpHandler := httphandlers.New(hUc, tUc)
	wsHandler := websockethandlers.New(wsUc)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"}

	s.router.Use(cors.New(config))

	//auth routes
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
	authRouts := s.router.Group("/", middleware.CheckAuthification())
	s.router.POST("/api/Account/SignIn", httpHandler.UserSignIn)
	s.router.POST("/api/Account/SignUp", httpHandler.UserSignUp)
	authRouts.POST("/api/Account/SignOut", httpHandler.UserSignOut)
	s.router.GET("/api/Account/RefreshTokens", httpHandler.RefreshTokens)
	authRouts.GET("/api/Account/Me", httpHandler.UserMyAccount)

<<<<<<< HEAD
	//websocket handlers
	authRouts.POST("/api/Room", httpHandler.CreateRoom)

=======
	//track route
	authRouts.GET("/api/Tracks", httpHandler.GetTracks)

	//room handler
	authRouts.POST("/api/Room", httpHandler.CreateRoom)

	//websocket handlers
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
	s.router.GET("/ws/:roomId/chat", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")

		wsHandler.ChatConnect(roomId)(ctx.Writer, ctx.Request)
	})

	s.router.GET("/ws/:roomId/file", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")

		wsHandler.FileConnect(roomId)(ctx.Writer, ctx.Request)
	})

<<<<<<< HEAD
=======
	s.router.GET("/ws/:roomId/track", func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")

		wsHandler.TrackConnect(roomId)(ctx.Writer, ctx.Request)
	})

>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
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
