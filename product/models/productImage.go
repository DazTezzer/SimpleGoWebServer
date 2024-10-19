package models

type ProductImage struct {
	Id    uint   `json:"id"  gorm:"column:id;primary_key"`
	Name  string `json:"name" gorm:"column:name"`
	ImageUrl string `json:"image" gorm:"column:image"`
}

func (ProductImage) TableName() string {
	return "productimage"
}
