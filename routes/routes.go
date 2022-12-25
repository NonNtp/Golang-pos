package routes

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func ServeRoutes(r *gin.Engine) {
	productController := controller.Product{}
	productGroup := r.Group("/api")

	productGroup.GET("/products", productController.FindAll)
	productGroup.GET("/product/:id", productController.FindById)
	productGroup.POST("/product", productController.Create)
	productGroup.PATCH("/product/:id", productController.Update)
	productGroup.DELETE("/product/:id", productController.Delete)

	categoryController := controller.Category{}
	categoryGroup := r.Group("/api")

	categoryGroup.GET("/categories", categoryController.FindAll)
	categoryGroup.GET("/category/:id", categoryController.FindById)
	categoryGroup.POST("/category", categoryController.Create)
	categoryGroup.PATCH("/category/:id", categoryController.Update)
	categoryGroup.DELETE("/category/:id", categoryController.Delete)

	orderController := controller.Order{}
	orderGroup := r.Group("/api")

	orderGroup.GET("/orders", orderController.FindAll)
	orderGroup.GET("/order/:id", orderController.FindById)
	orderGroup.POST("/order", orderController.Create)
}
