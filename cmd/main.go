package main

import (
	"log"
	"simple-forum/internals/configs"
	"simple-forum/internals/handlers/membership"
	membershipRepo "simple-forum/internals/repositories/membership"
	"simple-forum/pkg/internals"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	var (
		config *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internals/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatalf("Gagal menginisiaslisasi konfigurasi: %v", err)
	}

	config = configs.Get()

	db, err := internals.Connect(config.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}

	_ = membershipRepo.NewRepository(db)

	membershipHandler := membership.NewHandler(route)
	membershipHandler.RegisterRoutes()

	route.Run(config.Service.Port)
}
