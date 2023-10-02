package handler

import (
	"C_11S-practica_arquitectura/internal/products"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ProductsGetter interface para desacloparse de la logica del negocio
type ProductsGetter interface {
	GetByID(id int) (products.Product, error)
}

type ProductsCreator interface {
	ModifyByID(id int, product products.Product) (products.Product, error)
}

// ProductsHandler struct para poder inyectar la interface
type ProductsHandler struct {
	productsGetter  ProductsGetter
	productsCreator ProductsCreator
}

// NewProductsHandler contructor para inyectar la interface en el lugar donce usamos ese handler
func NewProductsHandler(getter ProductsGetter, creator ProductsCreator) *ProductsHandler {
	return &ProductsHandler{
		productsGetter:  getter,
		productsCreator: creator,
	}
}

func (ph *ProductsHandler) GetProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := ph.productsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", id)})
	}
	ctx.JSON(http.StatusOK, product)
}

func (ph *ProductsHandler) PutProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productRequest := products.Product{}
	err = ctx.BindJSON(&productRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	product, err := ph.productsCreator.ModifyByID(id, productRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, product)
}
