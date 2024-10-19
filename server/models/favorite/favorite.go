package favorite

type Favorite struct {
	Id         uint `json:"id"`
	CustomerID uint `json:"customer_id"`
	ProductID  uint `json:"product_id"`
}
