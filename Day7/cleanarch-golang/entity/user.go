package entity

type User struct {
	Id uint `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Firstname string `json:"firstname" gorm:"not_null"`
	Lastname string `json:"lastname" gorm:"not_null"`
	Email string `json:"email" gorm:"not_null"`
}

type CreateUserRequest struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
}

type UserResponse struct {
	Id       	uint    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
}