package models

type Cart struct {
	Id         uint `json:"id" gorm:"column:id;primary_key"`
	CustomerID uint `json:"customer_id" gorm:"column:customer_id"`
}

func (Cart) TableName() string {
	return "carts"
}