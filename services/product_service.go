package services

import (
	"RetailerAPI/database"
	"RetailerAPI/models"
)

// CreateProduct - Creates a new product
func CreateProduct(product *models.Product) error {
	db := database.GetDB() // Get the database instance
	return db.Create(product).Error
}

// GetAllProducts - Retrieves all products
func GetAllProducts(products *[]models.Product) error {
	db := database.GetDB()
	return db.Find(products).Error
}

// GetProductByID - Retrieves a product by ID
func GetProductByID(id string, product *models.Product) error {
	db := database.GetDB()
	return db.First(product, "id = ?", id).Error
}

// UpdateProduct - Updates an existing product by ID
func UpdateProduct(id string, updatedProduct *models.Product) error {
	db := database.GetDB()
	return db.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct).Error
}
