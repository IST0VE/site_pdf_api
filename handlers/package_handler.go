package handlers

import (
	"net/http"

	"github.com/IST0VE/site_pdf_api/models"
	"github.com/IST0VE/site_pdf_api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePackage(c *gin.Context) {
	var pkg models.Package
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreatePackage(&pkg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"package_id": result.InsertedID})
}

func GetPackage(c *gin.Context) {
	packageID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID"})
		return
	}

	pkg, err := services.GetPackageByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pkg)
}

func UpdatePackage(c *gin.Context) {
	packageID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID"})
		return
	}

	var pkg models.Package
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.UpdatePackage(objID, &pkg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAllPackages(c *gin.Context) {
	packages, err := services.GetAllPackages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, packages)
}

func DeletePackage(c *gin.Context) {
	packageID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(packageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID"})
		return
	}

	result, err := services.DeletePackage(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
