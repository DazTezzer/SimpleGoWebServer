package models

type Customer struct {
	Id           uint   `json:"id" gorm:"column:id;primary_key"`
	Username     string `json:"username" gorm:"column:username"`
	Email        string `json:"email" gorm:"column:email"`
	PasswordHash string `json:"password_hash" gorm:"column:password_hash"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at"`
}

func (Customer) TableName() string {
	return "customer"
}
