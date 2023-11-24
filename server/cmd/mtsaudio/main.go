package main

import (
	"context"
	"log"
	"mtsaudio/internal/config"
	"mtsaudio/internal/database"
	"mtsaudio/internal/tokens"
	"mtsaudio/internal/transport"
	"mtsaudio/internal/usecase/authusecase"
	"os"
	"os/signal"
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
	srv := transport.New(":80")
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer stop()

	srv.Run(ctx, authUsecase)
}
