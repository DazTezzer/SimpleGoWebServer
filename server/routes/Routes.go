package routes

import (
	"bebeziansback/cart"
	"bebeziansback/customer"
	_ "bebeziansback/docs"
	"bebeziansback/product"
	"bebeziansback/security"
	"bebeziansback/server/config"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.GetCORS()},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	productEndpoints := r.Group("/product")
	{
		productEndpoints.GET("/getAllCategory", product.GetAllCategory)
		productEndpoints.GET("/getPopProducts", product.GetPopProducts)
		productEndpoints.POST("/getProductsByCategory", product.GetProductsByCategory)
		productEndpoints.POST("/getProduct", product.GetProductById)
	}

	cartEndpoints := r.Group("/cart").Use(security.AuthMiddleware())
	{
		cartEndpoints.GET("/getCart", cart.GetCart)
		cartEndpoints.POST("/addToCart", cart.AddToCart)
	}

	customerEndpoints := r.Group("/customer")
	{
		customerEndpoints.POST("/register", customer.Register)
		customerEndpoints.POST("/login", customer.Login)
	}

	r.GET("/", GetTest)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// @Summary      Get Test
// @Description  Responds with the test answer as JSON.
// @Router       / [get]
// @Tags Server
func GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
