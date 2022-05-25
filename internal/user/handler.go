package user

import (
	"example/messages_api/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service service
}

func NewHandler(service service) handlers.Handler {
	return &handler{service: service}
}

func (h handler) Register(router *gin.Engine) {
	router.POST("/users/", h.CreateUser)
	router.GET("/users/", h.GetUsers)
	router.PATCH("/users/:id", h.EditUser)
	router.DELETE("/users/:id", h.DeleteUser)
}

func (h handler) CreateUser(context *gin.Context) {
	var newUser User
	if err := context.BindJSON(&newUser); err != nil {
		return
	}
	h.service.CreateUser(newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func (h handler) GetUsers(context *gin.Context) {
	users := h.service.GetUsers()
	context.IndentedJSON(http.StatusOK, users)
}

func (h handler) EditUser(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")
	err := h.service.EditUser(id, user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func (h handler) DeleteUser(context *gin.Context) {
	id := context.Param("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with id=%s was deleted", id)})
}
