package models

type Person struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Role     string `json:"role"`
}
