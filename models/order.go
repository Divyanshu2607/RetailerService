package models

type Order struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status"`
}
