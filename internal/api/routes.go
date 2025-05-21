package api

import (
	"todo-app/internal/api/handlers"
	"todo-app/internal/config"
	"todo-app/internal/repository"
	"todo-app/internal/services"
	"todo-app/pkg/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	jwtUtil := jwt.NewJWTUtil(cfg.JWTSecret)

	authService := services.NewAuthService(userRepo, jwtUtil)
	taskService := services.NewTaskService(taskRepo)

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	authGroup := router.Group("/").Use(handlers.AuthMiddleware(jwtUtil))
	{
		authGroup.GET("/tasks", taskHandler.GetTasks)
		authGroup.POST("/tasks", taskHandler.CreateTask)
		authGroup.PUT("/tasks/:id", taskHandler.UpdateTask)
		authGroup.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}

	return router
}