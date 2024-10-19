package models

type ProductCategory struct {
	ID   uint   `json:"id" gorm:"column:id;primary_key"`
	Name string `json:"name" gorm:"column:name"`
}

func (ProductCategory) TableName() string {
	return "productcategory"
}
