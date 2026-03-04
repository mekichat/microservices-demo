package api

import (
	"fmt"
	"microservices-demo/internals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


func StartServer(productService *service.ProductService) {
	
	r := gin.Default()

	r.POST("/products", func(ctx *gin.Context){

		var input struct {
			Name string `json:"name"`
			Price int `json:"price"`
		}

		if err := ctx.BindJSON(&input); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
        }

        product := productService.CreateProduct(input.Name, input.Price)
        ctx.JSON(http.StatusCreated, product)

	})

	r.GET("/products", func(ctx *gin.Context) {
        products := productService.ListProducts()
        ctx.JSON(http.StatusOK, products)
    })


	r.PUT("/products/:id", func(ctx *gin.Context) {
        var id uint
        fmt.Sscanf(ctx.Param("id"), "%d", &id)

        var input struct {
            Name string `json:"name"`
            Price int `json:"price"`
        }

        if err := ctx.BindJSON(&input); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
        }

        product, ok := productService.UpdateProduct(id, input.Name, input.Price)

        if !ok {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }
        ctx.JSON(http.StatusOK, product)
    })

	 r.DELETE("/products/:id", func(ctx *gin.Context) {
        var id uint
        fmt.Sscanf(ctx.Param("id"), "%d", &id)

        if !productService.DeleteProduct(id) {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }
        ctx.Status(http.StatusNoContent)
    })

	r.Run(":8080")
}
