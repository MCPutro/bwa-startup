package main

import (
	"bwa-startup/config"
	"bwa-startup/internal/database"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/routes"
	"bwa-startup/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.NewConfig()

	db, err := database.NewPostgre(cfg.DatabaseConf())

	if err != nil {
		fmt.Println("- error : ", err)
		return
	}

	repo := repository.NewRepoManagerImpl(cfg, db)
	services := service.NewServiceManagerImpl(cfg, repo)

	routes.RegisterRoute(r, services, repo, cfg)

	r.Run(":" + cfg.ServerConf().Port)

}
