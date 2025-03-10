package controllers

import (
	"RetailerAPI/models"
	"RetailerAPI/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProduct - Handles product creation
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateProduct(&product); err != nil { // Directly calling the service function
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProducts - Retrieves all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := services.GetAllProducts(&products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
