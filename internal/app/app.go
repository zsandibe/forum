package app

import (
	"forum/internal/config"
	"forum/internal/delivery"
	"forum/internal/repository"
	"forum/internal/server"
	"forum/internal/service"
	"forum/internal/storage"
	"forum/pkg"
	"log"
)

func Start() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	pkg.InfoLog.Println("Config loaded")

	db, err := storage.NewSqlite3(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	pkg.InfoLog.Println("DB connected")

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	delivery := delivery.NewHandler(service)

	server := new(server.Server)

	if err := server.Run(config.Port, delivery.Routes()); err != nil {
		log.Fatalf("Error running server: %s\n", err)
	}
}
