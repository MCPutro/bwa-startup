package main

import (
	"bwa-startup/config"
	"bwa-startup/internal/app"
	"bwa-startup/internal/database"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/routes"
	"bwa-startup/internal/service"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println("Failed load config. Error message :", err.Error())
		return
	}

	db, err := database.NewPostgre(cfg.DatabaseConf())

	if err != nil {
		log.Println("Failed create database connection. Error message :", err)
		return
	}

	log.Println("Failed load config. error message :", err.Error())
	server := app.NewServer()

	repo := repository.NewRepoManagerImpl(cfg, db)
	services := service.NewServiceManagerImpl(cfg, repo)

	routes.RegisterRoute(server, services, repo, cfg)

	server.Listen(":" + cfg.ServerConf().Port)

}
