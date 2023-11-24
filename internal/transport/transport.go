package transport

import (
	"context"
	"log"
	_ "mtsaudio/docs"
	"mtsaudio/internal/transport/handlers/httphandlers"
	"mtsaudio/internal/transport/handlers/middleware"
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

func (s *Server) Run(ctx context.Context, huc httphandlers.AuthUsecase) {
	//swagger route
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//auth routes
	httpHandler := httphandlers.New(huc)

	//user auth routes
	authRouts := s.router.Group("/", middleware.CheckAuthification())
	s.router.POST("/api/Account/SignIn", httpHandler.UserSignIn)
	s.router.POST("/api/Account/SignUp", httpHandler.UserSignUp)
	authRouts.POST("/api/Account/SignOut", httpHandler.UserSignOut)
	authRouts.GET("/api/Account/RefreshTokens", httpHandler.RefreshToken)
	authRouts.GET("/api/Account/Me", httpHandler.UserMyAccount)

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
