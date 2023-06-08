package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}
