package routes

import (
	"bebeziansback/cart"
	"bebeziansback/customer"
	_ "bebeziansback/docs"
	"bebeziansback/product"
	"bebeziansback/security"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	productEndpoints := r.Group("/product")
	{
		productEndpoints.GET("/getAllCategory", product.GetAllCategory)
		productEndpoints.GET("/getPopProducts", product.GetPopProducts)
		productEndpoints.POST("/getProductsByCategory", product.GetProductsByCategory)
		productEndpoints.POST("/getProduct", product.GetProductById)
	}

	cartEndpoints := r.Group("/cart").Use(security.AuthMiddleware())
	{
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
