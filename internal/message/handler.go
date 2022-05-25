package message

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

//func getUserByName(name string) (*user.User, error) {
//	var users []user.User
//	postgres.GetDB().Find(&users)
//
//	for _, user := range users {
//		if user.Name == name {
//			return &user, nil
//		}
//	}
//
//	return nil, errors.New("user not found")
//}

func (h *handler) Register(router *gin.Engine) {
	router.POST("/messages/", h.CreateMessage)
	router.GET("/messages/", h.GetMessages)
	router.PATCH("/messages/:id/", h.EditMessage)
	router.DELETE("/messages/:id/", h.DeleteMessage)
}

func (h *handler) CreateMessage(context *gin.Context) {
	var input CreateMessageInput
	if err := context.BindJSON(&input); err != nil {
		return
	}

	err := h.service.CreateMessage(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusCreated, input)
}

func (h *handler) GetMessages(context *gin.Context) {
	messages := h.service.GetMessages()
	context.IndentedJSON(http.StatusOK, messages)
}

func (h *handler) EditMessage(context *gin.Context) {
	var editInput EditMessageInput
	if err := context.ShouldBindJSON(&editInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := context.Param("id")

	err := h.service.EditMessage(id, editInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, editInput)
}

func (h *handler) DeleteMessage(context *gin.Context) {
	id := context.Param("id")
	err := h.service.repository.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Message with id=%s was deleted", id)})
}
