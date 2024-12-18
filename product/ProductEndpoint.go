package product

import (
	_ "bebeziansback/product/response"
	"bebeziansback/product/request"
	"bebeziansback/product/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Get All Categories
// @Description  Get all product categories
// @Produce      json
// @Success      200  {object}  response.ProductCategoryResponse  "OK"
// @Router       /api/product/getAllCategory [get]
// @Tags Product
func GetAllCategory(c *gin.Context) {

	categories, err := service.GetAllCategories()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, categories)
}

// @Summary      Get All Pop Products
// @Description  Get all Pop Products
// @Produce      json
// @Success      200  {object}  response.PopProductsResponse  "OK"
// @Router       /api/product/getPopProducts [get]
// @Tags Product
func GetPopProducts(c *gin.Context) {

	products, err := service.GetPopProducts()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary      Get All Products By Category
// @Description  Get products by category name
// @Accept 		 json
// @Produce      json
// @Param 		 body body request.ProductsByCategoryRequest true "Category Name"
// @Success      200   {object}  response.ProductsByCategoryResponse  "OK"
// @Router       /api/product/getProductsByCategory [post]
// @Tags Product
func GetProductsByCategory(c *gin.Context) {
	var req request.ProductsByCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, err := service.GetProductsByCategory(req.CategoryName)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary      Get Product
// @Description  Get Product by id
// @Accept 		 json
// @Produce      json
// @Param 		 body body request.ProductRequest true "Product Id"
// @Success      200   {object}  response.ProductResponse  "OK"
// @Router       /api/product/getProduct [post]
// @Tags Product
func GetProductById(c *gin.Context) {
	var req request.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := service.GetProductById(req.ProductId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}
