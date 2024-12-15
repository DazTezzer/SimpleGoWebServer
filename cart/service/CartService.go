package service

import (
	"bebeziansback/cart/models"
	"bebeziansback/cart/request"
	customer "bebeziansback/customer/models"
	"bebeziansback/server/config"
	"log"
	"time"

	"gorm.io/gorm"
)

func AddToCart(request request.AddToCartRequest, customerID uint) error {
	var customer customer.Customer
	if err := config.DB.Where("id = ?", customerID).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("ERROR: customer with ID %d does not exist", customerID)
			return err
		}
		return err
	}

	var cart models.Cart
	if err := config.DB.Where("customer_id = ?", customerID).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			cart = models.Cart{
				CustomerID: customerID,
			}
			if err := config.DB.Create(&cart).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	var cartItem models.CartItem
	if err := config.DB.Where("cart_id = ? AND product_id = ?", cart.Id, request.ProductID).First(&cartItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			cartItem = models.CartItem{
				CartID:    cart.Id,
				ProductID: request.ProductID,
				Quantity:  1,
				AddedAt:   time.Now().Format(time.RFC3339),
			}
			if err := config.DB.Create(&cartItem).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		cartItem.Quantity++
		if err := config.DB.Save(&cartItem).Error; err != nil {
			return err
		}
	}

	return nil
}
