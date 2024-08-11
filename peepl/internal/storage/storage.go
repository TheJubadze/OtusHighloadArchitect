package storage

import (
	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/model"
)

type Storage interface {
	AddUser(user User) error
	UpdateUser(user User) error
	DeleteUser(login string) error
	GetUser(login string) (User, error)
	ListUsers() ([]User, error)

	AddUserRole(role UserRole) error
	UpdateUserRole(role UserRole) error
	DeleteUserRole(id int) error
	GetUserRole(id int) (UserRole, error)
	ListUserRoles() ([]UserRole, error)

	AddCity(city City) error
	UpdateCity(city City) error
	DeleteCity(id int) error
	GetCity(id int) (City, error)
	ListCities() ([]City, error)
}
