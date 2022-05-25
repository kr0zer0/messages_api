package main

import (
	"example/messages_api/internal/message"
	messageDb "example/messages_api/internal/message/db"
	"example/messages_api/internal/user"
	userDb "example/messages_api/internal/user/db"
	"example/messages_api/pkg/client/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing the database
	postgres.Init()

	// setting up the router
	router := gin.Default()

	messageRepo := messageDb.NewRepository(postgres.GetDB())
	messageService := message.NewService(messageRepo)
	messageHandler := message.NewHandler(messageService)
	messageHandler.Register(router)

	userRepo := userDb.NewRepository(postgres.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	userHandler.Register(router)

	router.Run("localhost:8080")
}
