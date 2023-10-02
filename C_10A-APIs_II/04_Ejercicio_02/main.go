package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type product struct {
	ID          int     `json:"id"`
	Name        string  `json:"Name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type errorMessage struct {
	ErrorInfo string `json:"error_info"`
}

var products = []product{
	{
		ID:          1,
		Name:        "Cheese",
		Quantity:    5,
		CodeValue:   "AC1",
		IsPublished: true,
		Expiration:  "20/11/2024",
		Price:       100,
	},
	{
		ID:          2,
		Name:        "Onions",
		Quantity:    5,
		CodeValue:   "AC2",
		IsPublished: true,
		Expiration:  "20/11/2024",
		Price:       210,
	},
	{
		ID:          1,
		Name:        "Apples",
		Quantity:    5,
		CodeValue:   "AC4",
		IsPublished: true,
		Expiration:  "20/11/2024",
		Price:       400,
	},
}

func main() {

	router := gin.Default()

	productsGroup := router.Group("/products", mockAuthExample)

	productsGroup.GET("/search", searchProduct)

	router.Run()
}

func mockAuthExample(ctx *gin.Context) {
	fmt.Println("Entre a un mock de autenticaciÃ³n")
}

func searchProduct(ctx *gin.Context) {

	strPriceGT := ctx.Query("price")

	if strPriceGT == "" {
		ctx.JSON(http.StatusBadRequest, errorMessage{ErrorInfo: "Price is empty my friend"})
		return
	}

	priceGT, err := strconv.ParseFloat(strPriceGT, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorMessage{
			ErrorInfo: fmt.Sprintf("Price is bad formatted: %s", err.Error())})
		return
	}

	var productsResponse []product

	for _, value := range products {
		if value.Price > priceGT {
			productsResponse = append(productsResponse, value)
		}
	}

	ctx.JSON(http.StatusOK, productsResponse)
}
