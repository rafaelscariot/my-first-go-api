package controller

import (
	"github.com/gin-gonic/gin"
	"my-first-go-api/model"
	"my-first-go-api/usecase"
	"net/http"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{{
		ID:    1,
		Name:  "teste",
		Price: 1.0,
	}}

	ctx.JSON(http.StatusOK, products)
}
