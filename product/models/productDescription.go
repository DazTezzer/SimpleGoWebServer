package models

type ProductDescription struct {
	Id             uint   `json:"id"  gorm:"column:id;primary_key"`
	Color          string `json:"color" gorm:"column:color"`
	Material       string `json:"material" gorm:"column:material"`
	Gender         string `json:"gender" gorm:"column:gender"`
	AdditionalInfo string `json:"additional_info" gorm:"column:additional_info"`
}

func (ProductDescription) TableName() string {
	return "productdescription"
}
