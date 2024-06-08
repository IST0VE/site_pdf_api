package handlers

import (
	"net/http"

	"github.com/IST0VE/site_pdf_api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Создание временного почтового ящика
func CreateTempMailBox(c *gin.Context) {
	userID := c.MustGet("user_id").(primitive.ObjectID) // Получение из middleware или токена
	email, _ := services.GenerateUniqueEmail()          // Генерация уникального email
	mailbox := services.CreateMailbox(email, userID)
	c.JSON(http.StatusCreated, gin.H{"mailbox": mailbox})
}

// Получение всех сообщений из почтового ящика
func GetMessages(c *gin.Context) {
	mailboxID := c.Param("id")
	messages, err := services.GetMessages(mailboxID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving messages"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
