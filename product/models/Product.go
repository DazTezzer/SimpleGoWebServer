package models

type Product struct {
	Id            uint           `json:"id" gorm:"column:id;primary_key"`
	Name          string         `json:"name" gorm:"column:name"`
	ImageID       uint           `json:"-" gorm:"column:image_id"`
	Image         ProductImage   `gorm:"foreignKey:ImageID" json:"image_id"`
	Price         float64        `json:"price" gorm:"column:price"`
	DescriptionID uint           `json:"-" gorm:"column:description_id"`
	Description   ProductDescription `gorm:"foreignKey:DescriptionID" json:"description_id"`
	CategoryID    uint           `json:"-" gorm:"column:category_id"`
	Category      ProductCategory `gorm:"foreignKey:CategoryID" json:"category_id"`
}

func (Product) TableName() string {
	return "products"
}
