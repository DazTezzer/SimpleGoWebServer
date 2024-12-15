package service

import (
	customer "bebeziansback/customer/models"
	"bebeziansback/profile/response"
	"bebeziansback/server/config"
)

func ProfileGetInfo(customerID uint) (response.ProfileInfoResponse, error) {
	var customer customer.Customer
	if err := config.DB.
		Where("id = ?", customerID).
		First(&customer).Error; err != nil {
		return response.ProfileInfoResponse{}, err
	}
	profileInfo := response.ProfileInfoResponse{
		Name:  customer.Username,
		Email: customer.Email,
	}
	return profileInfo, nil
}
