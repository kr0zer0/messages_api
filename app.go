package main

import (
	"example/messages_api/controllers"
	"example/messages_api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing the database
	database.Init()

	// setting up the router
	router := gin.Default()

	// messages related routes
	router.POST("/messages/", controllers.CreateMessage)
	router.GET("/messages/", controllers.GetMessages)
	router.PATCH("/messages/:id/", controllers.EditMessage)
	router.DELETE("/messages/:id/", controllers.DeleteMessage)

	// users related routes
	router.POST("/users/", controllers.CreateUser)
	router.GET("/users/", controllers.GetUsers)
	router.PATCH("/users/:id", controllers.EditUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	router.Run("localhost:8080")
}
