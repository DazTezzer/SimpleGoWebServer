package routes

import (
	_ "GoProject/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GoProject Service
// @version         1.0
// @description     This is a sample server for GoProject
// @host            localhost:8081
// @BasePath        /
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", GetTest)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// @Summary      Get Test
// @Description  Responds with the test answer as JSON.
// @Router       / [get]
func GetTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
