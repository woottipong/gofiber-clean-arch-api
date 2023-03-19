package entity

type User struct {
	ID       string `json:"id"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `json:"-"`
}
