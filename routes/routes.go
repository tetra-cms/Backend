package routes

import (
	"github.com/gin-gonic/gin"

	"tetra-server/config"
	"tetra-server/handlers"
	"tetra-server/middleware"
	"tetra-server/models"
)

func Setup(
	router *gin.Engine,
	cfg *config.Config,

	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	clientHandler *handlers.ClientHandler,
	categoryHandler *handlers.CategoryHandler,
	productHandler *handlers.ProductHandler,
) {

	api := router.Group("/api")

	public := api.Group("/")

	public.POST("/auth/register", authHandler.Register)
	public.POST("/auth/login", authHandler.Login)
	public.POST("/auth/refresh", authHandler.Refresh)

	public.GET("/products", productHandler.GetAll)
	public.GET("/products/:id", productHandler.GetByID)

	public.GET("/categories", categoryHandler.GetAll)
	public.GET("/categories/:id", categoryHandler.GetByID)

	authenticated := api.Group("/")

	authenticated.Use(middleware.AuthMiddleware(cfg))

	authenticated.GET("/auth/me", authHandler.Me)

	authenticated.GET("/clients", clientHandler.GetAll)
	authenticated.GET("/clients/:id", clientHandler.GetByID)
	authenticated.POST("/clients", clientHandler.Create)
	authenticated.PUT("/clients/:id", clientHandler.Update)
	authenticated.DELETE("/clients/:id", clientHandler.Delete)

	authenticated.GET("/users/profile", userHandler.Profile)

	employee := api.Group("/employee")

	employee.Use(middleware.AuthMiddleware(cfg))
	employee.Use(middleware.RequireRole(
		models.RoleEmployee,
		models.RoleAdmin,
	))

	employee.POST("/products", productHandler.Create)
	employee.PUT("/products/:id", productHandler.Update)

	admin := api.Group("/admin")

	admin.Use(middleware.AuthMiddleware(cfg))
	admin.Use(middleware.RequireRole(models.RoleAdmin))

	// Users
	admin.GET("/users", userHandler.GetAll)
	admin.GET("/users/:id", userHandler.GetByID)
	admin.PUT("/users/:id", userHandler.Update)
	admin.DELETE("/users/:id", userHandler.Delete)

	// Categories
	admin.POST("/categories", categoryHandler.Create)
	admin.PUT("/categories/:id", categoryHandler.Update)
	admin.DELETE("/categories/:id", categoryHandler.Delete)

	// Products
	admin.DELETE("/products/:id", productHandler.Delete)
}
