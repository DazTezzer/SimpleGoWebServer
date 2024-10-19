package response

type ProductResponsePart struct {
	ProductID       uint   `json:"productId"`
	Name            string `json:"name"`
	ImageURL        string `json:"image_url"`
	Price           string `json:"price"`
	DescriptionInfo string `json:"description_info"`
}
