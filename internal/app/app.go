package app

import (
	"casbin-go_gin/internal/app/handler"
	"casbin-go_gin/internal/repository"
	"casbin-go_gin/internal/services"
	"context"
	"log"
)

func Init(port string) {
	db, adapter, err := repository.NewPostgresDB(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	service := services.NewService(repos)
	handlers := handler.NewHandler(service, adapter)

	srv := new(Server)

	if err = srv.Run(":"+port, handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
