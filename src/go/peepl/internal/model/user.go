package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"password" db:"password"`
	Firstname string    `json:"firstname" db:"firstname"`
	Lastname  string    `json:"lastname" db:"lastname"`
	Birthdate time.Time `json:"birthdate" db:"birthdate"`
	Sex       string    `json:"sex" db:"sex"`
	Interests *string   `json:"interests,omitempty" db:"interests"` // Interests can be null, so using a pointer
	CityID    int       `json:"city_id" db:"city_id"`
	RoleID    int       `json:"role_id" db:"role_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
