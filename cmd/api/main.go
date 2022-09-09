package main

import (
	"ecommerce-api/internal/database"

	"github.com/gin-gonic/gin"

	"ecommerce-api/internal/middlewares"
	"ecommerce-api/internal/product/usecase"
	userUc "ecommerce-api/internal/user/usecase"
)

func main() {
	r := gin.Default()

	db := database.NewDatabaseConn()
	productUsecase := usecase.NewProductUseCase(db)
	userUcs := userUc.NewUserUseCase(db)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/products", middlewares.WithAuth(userUcs), productUsecase.CreateProduct)
	r.GET("/products", productUsecase.GetProducts)
	r.GET("/products/:id", productUsecase.GetProduct)
	r.PUT("/products/:id", middlewares.WithAuth(userUcs), productUsecase.UpdateProduct)
	r.DELETE("/products/:id", middlewares.WithAuth(userUcs), productUsecase.DeleteProduct)
	r.GET("/products/search?name=:name", productUsecase.SearchProduct)
	r.GET("/products/category/:category", productUsecase.GetProductByCategory)
	r.GET("/products/searchByPrice", productUsecase.GetProductByPrice)
	r.GET("/products/searchByStock", productUsecase.GetProductByStock)
	r.POST("/categories", middlewares.WithAuth(userUcs), productUsecase.CreateCategory)
	r.GET("/categories", productUsecase.GetCategory)
	r.GET("/categories/:id", productUsecase.GetCategoryByID)
	r.PUT("/categories/:id", middlewares.WithAuth(userUcs), productUsecase.UpdateCategory)
	r.DELETE("/categories/:id", middlewares.WithAuth(userUcs), productUsecase.DeleteCategory)

	r.POST("/products/buy/:id", middlewares.WithAuth(userUcs), productUsecase.BuyProduct)

	r.POST("/register", userUcs.Register)
	r.POST("/login", userUcs.Login)

	r.Run(":5000") // listen and serve on
}
