package main

import (
	"context"
	"log"
	"mtsaudio/internal/config"
	"mtsaudio/internal/database"
	"mtsaudio/internal/tokens"
	"mtsaudio/internal/transport"
	"mtsaudio/internal/usecase/authusecase"
	websocketusecase "mtsaudio/internal/usecase/roomusecase"
<<<<<<< HEAD
	"os"
	"os/signal"
=======
	"mtsaudio/internal/usecase/trackusecase"
	"os"
	"os/signal"
	"path/filepath"
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
	"syscall"
)

// @title           MTSAudio
// @version         1.0
// @description     Server for mtsaudio
// @termsOfService  http://swagger.io/terms/

// @contact.name   Alexander Soldatov
// @contact.email  soldatovalex207z@gmail.com

// @host      localhost:80
// @BasePath  /

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	cfg := config.Init()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("succesfully connect to database")

	tokens.InitBlackList()

	authUsecase := authusecase.New(db)
<<<<<<< HEAD
=======
	wsUsecase := websocketusecase.New(db)
	dirPath, err := filepath.Abs("./static/")
	trackUsecase := trackusecase.New(dirPath)

>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
	srv := transport.New(":80")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer stop()

<<<<<<< HEAD
	wsUsecase := websocketusecase.New(db)

	srv.Run(ctx, authUsecase, wsUsecase)
=======
	srv.Run(ctx, authUsecase, wsUsecase, trackUsecase)
>>>>>>> 7b9fc1984b09754619ee44ea4dd07a95ad8762c2
}
