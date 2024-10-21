package request

type AddToCartRequest struct {
	CustomerID uint   `json:"clientId"`
	ProductID  uint   `json:"productId"`
	Size       string `json:"size"`
}
