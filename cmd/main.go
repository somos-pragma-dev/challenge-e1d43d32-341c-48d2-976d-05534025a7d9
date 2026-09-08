package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"internal/handler"
	"internal/repository"
	"internal/service"
	"pkg/validation"
)

func main() {
	router := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err!= nil {
		log.Fatal("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, validation.NewValidator())
	userHandler := handler.NewUserHandler(userService)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	if err := router.Run(":8080"); err!= nil {
		log.Fatal("failed to start server")
	}
}