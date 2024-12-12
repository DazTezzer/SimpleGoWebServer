package service

import (
	"bebeziansback/customer/models"
	"bebeziansback/customer/request"
	"bebeziansback/customer/response"
	"bebeziansback/security"
	"bebeziansback/server/config"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterCustomer(customerRequest request.CustomerRequest) error {
	err := checkCustomer(customerRequest.Email)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(customerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	customer := createCustomer(customerRequest, hash)
	if err := config.DB.Save(&customer).Error; err != nil {
		return err
	}
	return nil
}

func LoginCustomer(customerRequest request.CustomerRequest) (response.TokenResponse, error) {
	var foundCustomer models.Customer

	if err := config.DB.Where("email = ?", customerRequest.Email).First(&foundCustomer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.TokenResponse{}, err
		}
		return response.TokenResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundCustomer.PasswordHash), []byte(customerRequest.Password)); err != nil {
		return response.TokenResponse{}, err
	}
	token, err := security.GenerateToken(customerRequest.Username)
	if err != nil {
		return response.TokenResponse{}, err
	}

	return response.TokenResponse{Token: token}, nil
}

func createCustomer(customerRequest request.CustomerRequest, hash []byte) models.Customer {
	return models.Customer{
		Username:     customerRequest.Username,
		PasswordHash: string(hash),
		Email:        customerRequest.Email,
	}
}

func checkCustomer(customerEmail string) error {
	var customer models.Customer
	if err := config.DB.Where("email = ?", customerEmail).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	log.Printf("ERROR: such a user already exists with email: %s", customerEmail)
	return fmt.Errorf("пользователь с таким email уже существует")
}
