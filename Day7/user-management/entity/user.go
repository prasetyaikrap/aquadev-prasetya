package entity

type User struct {
	Id 				uint64 	`json:"id" gorm:"primary_key;auto_increment;not_null"`
	Firstname string 	`json:"firstname" gorm:"not_null"`
	Lastname 	string 	`json:"lastname" gorm:"not_null"`
	Email 		string 	`json:"email" gorm:"not_null"`
	Gender 		string 	`json:"gender" gorm:"not_null"`
}

type CreateUserRequest struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Gender 		string 		`json:"gender"`
}

type UserResponse struct {
	Id       	uint64    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Gender 		string 		`json:"gender"`
}