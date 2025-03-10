package services

import (
	"RetailerAPI/database"
	"RetailerAPI/models"
)

// CreateOrder - Creates a new order
func CreateOrder(order *models.Order) error {
	db := database.GetDB() // Get the database instance
	return db.Create(order).Error
}

// GetOrderByID - Retrieves an order by ID
func GetOrderByID(id string, order *models.Order) error {
	db := database.GetDB()
	return db.First(order, "id = ?", id).Error
}
