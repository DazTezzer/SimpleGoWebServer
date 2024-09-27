package product

type Product struct {
	Id            uint    `json:"id"`
	Name          string  `json:"name"`
	ImageID       uint    `json:"image_id"`
	Price         float64 `json:"price"`
	DescriptionID uint    `json:"description_id"`
	CategoryID    uint    `json:"category_id"`
}
