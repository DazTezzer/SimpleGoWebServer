package profile

import (
	"bebeziansback/profile/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Get Profile Info
// @Description  Get Profile Info by customer
// @Accept 		 json
// @Produce      json
// @Security     BearerAuth
// @Success      200 "OK"
// @Router       /api/profile/info [get]
// @Tags Profile
func Info(c *gin.Context) {
	customerIDStr := c.MustGet("customerID").(string)
	customerID64, err := strconv.ParseUint(customerIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор клиента:"})
		return
	}
	customerID := uint(customerID64)
	info, err := service.ProfileGetInfo(customerID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, info)
}
