package routes

import (
	"RetailerAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/product", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)

	r.POST("/order", controllers.CreateOrder)
	r.GET("/order/:id", controllers.GetOrderByID)

	return r
}
