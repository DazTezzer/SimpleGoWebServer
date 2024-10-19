package response

type ProductCategoryResponse struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Name string `json:"name" gorm:"column:name"`
}
