package service

import (
	"bebeziansback/cart/models"
	"bebeziansback/cart/request"
	"bebeziansback/cart/response"
	customer "bebeziansback/customer/models"
	product "bebeziansback/product/models"
	"bebeziansback/server/config"
	"fmt"
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
	if err := config.DB.Where("cart_id = ? AND product_id = ? AND size = ?", cart.Id, request.ProductID, request.Size).First(&cartItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			cartItem = models.CartItem{
				CartID:    cart.Id,
				ProductID: request.ProductID,
				Size:      request.Size,
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

func GetCart(customerID uint) (response.CartResponse, error) {
	var cart models.Cart
	if err := config.DB.
		Where("customer_id = ?", customerID).
		First(&cart).Error; err != nil {
		return response.CartResponse{}, err
	}

	var cartItems []models.CartItem
	if err := config.DB.
		Where("cart_id = ?", cart.Id).
		Find(&cartItems).Error; err != nil {
		return response.CartResponse{}, err
	}

	var cartProducts []models.CartProduct
	var totalPrice float64
	for _, cartItem := range cartItems {
		var product product.Product
		if err := config.DB.
			Preload("Image").
			Preload("Description").
			Preload("Category").
			Where("id = ?", cartItem.ProductID).
			First(&product).Error; err != nil {
			return response.CartResponse{}, err
		}

		cartProduct := models.CartProduct{
			CartItemId:      cartItem.Id,
			ProductID:       cartItem.ProductID,
			Name:            product.Name,
			ImageURL:        product.Image.ImageUrl,
			Price:           fmt.Sprintf("%.2f", product.Price),
			DescriptionInfo: product.Description.AdditionalInfo,
			Size:            cartItem.Size,
			Quantity:        cartItem.Quantity,
			SumPrice:        fmt.Sprintf("%.2f", product.Price*float64(cartItem.Quantity)),
		}
		cartProducts = append(cartProducts, cartProduct)
		totalPrice += product.Price * float64(cartItem.Quantity)
	}

	cartResponse := response.CartResponse{
		CartProduct: cartProducts,
		TotalPrice:  fmt.Sprintf("%.2f", totalPrice),
	}
	return cartResponse, nil
}
