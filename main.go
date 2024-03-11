package main

import (
	"bwa-startup/config"
	"bwa-startup/internal/app"
	"bwa-startup/internal/database"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/routes"
	"bwa-startup/internal/service"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println("Failed load config. error message :", err.Error())
		return
	}

	db, err := database.NewPostgre(cfg.DatabaseConf())

	if err != nil {
		fmt.Println("- error : ", err)
		return
	}

	server := app.NewServer()

	repo := repository.NewRepoManagerImpl(cfg, db)
	services := service.NewServiceManagerImpl(cfg, repo)

	routes.RegisterRoute(server, services, repo, cfg)

	server.Listen(":" + cfg.ServerConf().Port)

}
