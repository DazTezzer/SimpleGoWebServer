package routes

import (
	"bebeziansback/cart"
	"bebeziansback/customer"
	_ "bebeziansback/docs"
	"bebeziansback/product"
	"bebeziansback/profile"
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
		AllowOrigins:     config.GetCORS(),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	apiGroup := r.Group("/api")

	productEndpoints := apiGroup.Group("/product")
	{
		productEndpoints.GET("/getAllCategory", product.GetAllCategory)
		productEndpoints.GET("/getPopProducts", product.GetPopProducts)
		productEndpoints.POST("/getProductsByCategory", product.GetProductsByCategory)
		productEndpoints.POST("/getProduct", product.GetProductById)
	}

	cartEndpoints := apiGroup.Group("/cart").Use(security.AuthMiddleware())
	{
		cartEndpoints.GET("/getCart", cart.GetCart)
		cartEndpoints.POST("/addToCart", cart.AddToCart)
	}

	customerEndpoints := apiGroup.Group("/customer")
	{
		customerEndpoints.POST("/register", customer.Register)
		customerEndpoints.POST("/login", customer.Login)
	}

	profileEndpoints := apiGroup.Group("/profile").Use(security.AuthMiddleware())
	{
		profileEndpoints.GET("/info", profile.Info)
	}

	r.GET("/api/", GetTest)
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// @Summary      Get Test
// @Description  Responds with the test answer as JSON.
// @Router       /api/ [get]
// @Tags Server
func GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
