package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"tetra-server/config"
	"tetra-server/database"
	"tetra-server/handlers"
	"tetra-server/middleware"
	"tetra-server/models"
	"tetra-server/routes"
)

func main() {
	cfg := config.Load()

	if err := database.Connect(cfg); err != nil {
		log.Fatal(err)
	}

	if err := database.DB.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.Category{},
		&models.Product{},
		&models.RefreshToken{},
	); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(middleware.Logger())

	authHandler := handlers.NewAuthHandler(cfg)
	userHandler := handlers.NewUserHandler()
	clientHandler := handlers.NewClientHandler()
	categoryHandler := handlers.NewCategoryHandler()
	productHandler := handlers.NewProductHandler()

	routes.Setup(
		router,
		cfg,

		authHandler,
		userHandler,
		clientHandler,
		categoryHandler,
		productHandler,
	)

	log.Printf("Server started on :%s", cfg.Port)

	router.Run(":" + cfg.Port)
}
