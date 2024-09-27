package review

type Review struct {
	Id        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	CustomerID uint   `json:"customer_id"`
	Rating    uint   `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
	OrderID   uint   `json:"order_id"`
}