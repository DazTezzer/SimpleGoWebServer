package models

type CartItem struct {
	CartID    uint   `json:"cart_id" gorm:"column:cart_id;primary_key"`
	ProductID uint   `json:"product_id" gorm:"column:product_id"`
	Quantity  uint   `json:"quantity" gorm:"column:quantity"`
	AddedAt   string `json:"added_at" gorm:"column:added_at"`
}

func (CartItem) TableName() string {
	return "cartitems"
}
