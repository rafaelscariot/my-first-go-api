package main

import (
	"github.com/gin-gonic/gin"
	"my-first-go-api/controller"
	"my-first-go-api/usecase"
)

func main() {
	ProductUsecase := usecase.NewProductUseCase()

	ProductController := controller.NewProductController(ProductUsecase)

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8080")
}
