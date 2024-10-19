package cartItem

type CartItem struct {
	CartID    uint   `json:"cart_id"`
	ProductID uint   `json:"product_id"`
	Quantity  uint   `json:"quantity"`
	AddedAt   string `json:"added_at"`
}
