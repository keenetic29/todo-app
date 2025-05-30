package api

import (
	_ "todo-app/docs"
	"todo-app/internal/api/handlers"
	"todo-app/internal/config"
	"todo-app/internal/repository"
	"todo-app/internal/services"
	"todo-app/pkg/jwt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	categoryRepo := repository.NewCategoryRepository(db) 

	jwtUtil := jwt.NewJWTUtil(cfg.JWTSecret)

	authService := services.NewAuthService(userRepo, jwtUtil)
	taskService := services.NewTaskService(taskRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo) 
	

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)
	categoryHandler := handlers.NewCategoryHandler(categoryService) 

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	authGroup := router.Group("/").Use(handlers.AuthMiddleware(jwtUtil))
	{
		authGroup.GET("/tasks", taskHandler.GetTasks)
		authGroup.GET("/tasks/:id", taskHandler.GetTaskByID)
		authGroup.POST("/tasks", taskHandler.CreateTask)
		authGroup.PUT("/tasks/:id", taskHandler.UpdateTask)
		authGroup.DELETE("/tasks/:id", taskHandler.DeleteTask)

		authGroup.POST("/categories", categoryHandler.CreateCategory)
        authGroup.GET("/categories", categoryHandler.GetCategories)
        authGroup.DELETE("/categories/:id", categoryHandler.DeleteCategory)

		authGroup.PATCH("/tasks/:id/category", taskHandler.UpdateTaskCategory)
	}

	return router
}