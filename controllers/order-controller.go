package controllers

import (
	"RetailerAPI/models"
	"RetailerAPI/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrder - Handles order creation
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrderByID - Retrieves an order by ID
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := services.GetOrderByID(id, &order); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
