package models

type User struct {
	Username string  `json:"username" gorm:"column:username"`
	Password string	 `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}