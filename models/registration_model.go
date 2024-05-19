package models
type RegistrationModel struct {
	ID     int64 `json:"id"`
	Name   string `json:"name"`
	Email string    `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
	Password string `json:"password"`
}