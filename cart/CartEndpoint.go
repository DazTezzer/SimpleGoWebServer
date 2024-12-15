package cart

import (
	"bebeziansback/cart/request"
	"bebeziansback/cart/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Add To Cart
// @Description  Add Product To Cart
// @Accept 		 json
// @Produce      json
// @Param 		 body body request.AddToCartRequest true "Cart Request"
// @Security     BearerAuth
// @Success      200 "OK"
// @Router       /cart/addToCart [post]
// @Tags Cart
func AddToCart(c *gin.Context) {
	var addRequest request.AddToCartRequest
	customerIDStr := c.MustGet("customerID").(string)
	customerID64, err := strconv.ParseUint(customerIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор клиента:"})
		return
	}
	customerID := uint(customerID64)
	if err := c.ShouldBindJSON(&addRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := service.AddToCart(addRequest, customerID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
