package main

import (
	"bwa-startup/config"
	"bwa-startup/internal/app"
	"bwa-startup/internal/database"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/routes"
	"bwa-startup/internal/service"
	"fmt"
)

func main() {
	cfg := config.NewConfig()

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
