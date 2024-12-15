package request

type AddToCartRequest struct {
	ProductID uint   `json:"productId"`
	Size      string `json:"size"`
}
