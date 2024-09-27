package order

type Order struct {
	Id          uint   `json:"id"`
	CustomerID  uint   `json:"customer_id"`
	OrderDate   string `json:"order_date"` // лучше использовать time.Time
	TotalAmount float64 `json:"total_amount"`
	Status      string `json:"status"`
}