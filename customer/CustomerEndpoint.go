package customer

import (
	"bebeziansback/customer/request"
	_ "bebeziansback/customer/response"
	"bebeziansback/customer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Register Customer
// @Description  Register Customer with name and password
// @Accept 		 json
// @Produce      json
// @Param 		 body body request.CustomerRegisterRequest true "Customer"
// @Success      200  "Customer created successfully"
// @Router       /customer/register [post]
// @Tags Customer
func Register(c *gin.Context) {
	var CustomerRequest request.CustomerRegisterRequest
	if err := c.ShouldBindJSON(&CustomerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := service.RegisterCustomer(CustomerRequest); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary      Login Customer
// @Description  Login Customer with name and password
// @Accept 		 json
// @Produce      json
// @Param 		 body body request.CustomerLoginRequest true "Customer"
// @Success 	 200 {object} response.TokenResponse "Customer logged in successfully"
// @Router       /customer/login [post]
// @Tags Customer
func Login(c *gin.Context) {
	var CustomerRequest request.CustomerLoginRequest
	if err := c.ShouldBindJSON(&CustomerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := service.LoginCustomer(CustomerRequest)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, token)
}