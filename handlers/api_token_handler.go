package handlers

import (
	"net/http"

	"github.com/IST0VE/site_pdf_api/models"
	"github.com/IST0VE/site_pdf_api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAPIToken(c *gin.Context) {
	var apiToken models.APIToken
	if err := c.ShouldBindJSON(&apiToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateAPIToken(&apiToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token_id": result.InsertedID})
}

func GetAPIToken(c *gin.Context) {
	tokenID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(tokenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
		return
	}

	apiToken, err := services.GetAPITokenByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, apiToken)
}

func GetAPITokensByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	apiTokens, err := services.GetAPITokensByUserID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, apiTokens)
}

func UpdateAPIToken(c *gin.Context) {
	tokenID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(tokenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
		return
	}

	var apiToken models.APIToken
	if err := c.ShouldBindJSON(&apiToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdateAPIToken(objID, &apiToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteAPIToken(c *gin.Context) {
	tokenID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(tokenID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
		return
	}

	result, err := services.DeleteAPIToken(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
