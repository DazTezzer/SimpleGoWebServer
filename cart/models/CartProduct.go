package models

type CartProduct struct {
	CartItemId      uint   `json:"cartItemId"`
	ProductID       uint   `json:"productId"`
	Name            string `json:"name"`
	ImageURL        string `json:"image_url"`
	Price           string `json:"price"`
	DescriptionInfo string `json:"description_info"`
	Size            string `json:"size"`
	Quantity        uint   `json:"quantity"`
	SumPrice        string `json:"sum_price"`
}
