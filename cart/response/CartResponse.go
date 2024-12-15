package response

import "bebeziansback/cart/models"

type CartResponse struct {
	CartProduct []models.CartProduct `json:"cart_product"`
	TotalPrice  string                 `json:"total_price"`
}
