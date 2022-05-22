package controllers

import (
	"example/messages_api/database"
	"example/messages_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser
// POST /users/
func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	database.GetDB().Create(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// GetUsers
// GET /users/
func GetUsers(c *gin.Context) {
	var users []models.User
	database.GetDB().Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

// EditUser
// PATCH /users/:id/
func EditUser(c *gin.Context) {
	var user models.User

	if err := database.GetDB().Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.GetDB().Model(&user).Updates(user)
	c.IndentedJSON(http.StatusOK, user)
}

// DeleteUser
// DELETE /users/:id/
func DeleteUser(c *gin.Context) {
	var user models.User

	if err := database.GetDB().Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.GetDB().Delete(&user)
	c.IndentedJSON(http.StatusOK, user)
}
