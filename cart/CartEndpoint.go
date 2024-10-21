package cart

import (
	"bebeziansback/cart/request"
	"bebeziansback/cart/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Add To Cart
// @Description  Add Product To Cart
// @Accept 		 json
// @Produce      json
// @Param 		 body body request.AddToCartRequest true "Cart Request"
// @Success      200 "OK"
// @Router       /cart/addToCart [post]
// @Tags Cart
func AddToCart(c *gin.Context) {
	var addRequest request.AddToCartRequest
	if err := c.ShouldBindJSON(&addRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := service.AddToCart(addRequest); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
