package controllers

import (
	"errors"
	"example/messages_api/database"
	"example/messages_api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Helper function to get models.User object by its name
func getUserByName(name string) (*models.User, error) {
	var users []models.User
	database.GetDB().Find(&users)

	for _, user := range users {
		if user.Name == name {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

// CreateMessage
// POST /messages/
func CreateMessage(c *gin.Context) {

	var input models.CreateMessageInput

	if err := c.BindJSON(&input); err != nil {
		return
	}

	sender, err := getUserByName(input.SenderName)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("User with name '%s' not found", input.SenderName)})
		return
	}

	receiver, err := getUserByName(input.ReceiverName)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("User with name '%s' not found", input.ReceiverName)})
		return
	}

	newMessage := models.Message{SenderID: sender.ID, ReceiverID: receiver.ID, MessageBody: input.MessageBody}
	database.GetDB().Create(&newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}

// GetMessages
// GET /messages/
func GetMessages(c *gin.Context) {
	var messages []models.Message
	database.GetDB().Find(&messages)

	c.IndentedJSON(http.StatusOK, messages)
}

// EditMessage
// PATCH /messages/:id/
func EditMessage(c *gin.Context) {
	var editInput models.EditMessageInput
	var message models.Message
	if err := database.GetDB().Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&editInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if editInput.MessageBody != "" {
		message.MessageBody = editInput.MessageBody
	}

	database.GetDB().Model(&message).Updates(message)
	c.IndentedJSON(http.StatusOK, message)
}

// DeleteMessage
// DELETE /messages/:id/
func DeleteMessage(c *gin.Context) {
	var message models.Message
	if err := database.GetDB().Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.GetDB().Delete(&message)
	c.IndentedJSON(http.StatusOK, message)
}
