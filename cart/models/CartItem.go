package models

type CartItem struct {
	Id        uint   `json:"id" gorm:"column:id;primary_key"`
	CartID    uint   `json:"cart_id" gorm:"column:cart_id;primary_key"`
	ProductID uint   `json:"product_id" gorm:"column:product_id"`
	Size      string `json:"size" gorm:"column:size"`
	Quantity  uint   `json:"quantity" gorm:"column:quantity"`
	AddedAt   string `json:"added_at" gorm:"column:added_at"`
}

func (CartItem) TableName() string {
	return "cartitems"
}
