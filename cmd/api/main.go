package main

import (
	"ecommerce-api/internal/database"

	"github.com/gin-gonic/gin"

	userUc "ecommerce-api/internal/user/usecase"
)

func main() {
	r := gin.Default()

	db := database.NewDatabaseConn()
	userUcs := userUc.NewUserUseCase(db)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", userUcs.Register)
	r.POST("/login", userUcs.Login)

	r.Run(":5000") // listen and serve on
}
