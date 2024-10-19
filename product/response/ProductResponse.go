package response

import "bebeziansback/product/models"

type ProductResponse struct {
	ProductID   uint                      `json:"productId"`
	Name        string                    `json:"name"`
	ImageURL    string                    `json:"image_url"`
	Price       string                    `json:"price"`
	Description models.ProductDescription `json:"description"`
}
