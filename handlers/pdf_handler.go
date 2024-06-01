package handlers

import (
	"net/http"

	"github.com/IST0VE/site_pdf_api/services"
	"github.com/gin-gonic/gin"
)

func GeneratePDF(c *gin.Context) {
	var request struct {
		HTMLContent string `json:"html_content" binding:"required"`
		APIToken    string `json:"api_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check and decrement API token usage
	err := services.DecrementAPITokenUsage(request.APIToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired API token"})
		return
	}

	// Generate PDF
	pdfBytes, err := services.GeneratePDF(request.HTMLContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set response headers
	c.Header("Content-Disposition", "attachment; filename=output.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
