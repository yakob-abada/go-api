package entity

type User struct {
	Id        int8   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
}
