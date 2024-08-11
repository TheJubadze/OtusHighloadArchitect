package storage

import (
	"fmt"

	. "github.com/TheJubadze/OtusHighloadArchitect/peepl/internal/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose"
)

type SQLStorage struct {
	db *sqlx.DB
}

func NewSqlStorage(dsn string) (*SQLStorage, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &SQLStorage{db: db}, nil
}

func (s *SQLStorage) Migrate(migrationsDir string) (err error) {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("cannot set dialect: %w", err)
	}
	if err := goose.Up(s.db.DB, migrationsDir); err != nil {
		return fmt.Errorf("cannot do up migration: %w", err)
	}

	return nil
}

func (s *SQLStorage) Close() error {
	return s.db.Close()
}

func (s *SQLStorage) AddUser(user User) error {
	query := `INSERT INTO users (login, password, firstname, lastname, birthdate, sex, interests, city_id, role_id)
			  VALUES (:login, :password, :firstname, :lastname, :birthdate, :sex, :interests, :city_id, :role_id)`
	_, err := s.db.NamedExec(query, &user)
	return err
}

func (s *SQLStorage) UpdateUser(user User) error {
	query := `UPDATE users SET login = :login, password = :password, firstname = :firstname, lastname = :lastname,
			  birthdate = :birthdate, sex = :sex, interests = :interests, city_id = :city_id, role_id = :role_id 
			  WHERE id = :id`
	_, err := s.db.NamedExec(query, &user)
	return err
}

func (s *SQLStorage) DeleteUser(login string) error {
	query := `DELETE FROM users WHERE login = $1`
	_, err := s.db.Exec(query, login)
	return err
}

func (s *SQLStorage) GetUser(login string) (User, error) {
	var user User
	query := `SELECT * FROM users WHERE login = $1`
	err := s.db.Get(&user, query, login)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *SQLStorage) ListUsers() ([]User, error) {
	var users []User
	query := `SELECT * FROM users`
	err := s.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *SQLStorage) AddUserRole(role UserRole) error {
	query := `INSERT INTO user_roles (role) VALUES (:role)`
	_, err := s.db.NamedExec(query, &role)
	return err
}

func (s *SQLStorage) UpdateUserRole(role UserRole) error {
	query := `UPDATE user_roles SET role = :role WHERE id = :id`
	_, err := s.db.NamedExec(query, &role)
	return err
}

func (s *SQLStorage) DeleteUserRole(id int) error {
	query := `DELETE FROM user_roles WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *SQLStorage) GetUserRole(id int) (UserRole, error) {
	var role UserRole
	query := `SELECT * FROM user_roles WHERE id = $1`
	err := s.db.Get(&role, query, id)
	if err != nil {
		return UserRole{}, err
	}
	return role, nil
}

func (s *SQLStorage) ListUserRoles() ([]UserRole, error) {
	var roles []UserRole
	query := `SELECT * FROM user_roles`
	err := s.db.Select(&roles, query)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *SQLStorage) AddCity(city City) error {
	query := `INSERT INTO cities (name) VALUES (:name)`
	_, err := s.db.NamedExec(query, &city)
	return err
}

func (s *SQLStorage) UpdateCity(city City) error {
	query := `UPDATE cities SET name = :name WHERE id = :id`
	_, err := s.db.NamedExec(query, &city)
	return err
}

func (s *SQLStorage) DeleteCity(id int) error {
	query := `DELETE FROM cities WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *SQLStorage) GetCity(id int) (City, error) {
	var city City
	query := `SELECT * FROM cities WHERE id = $1`
	err := s.db.Get(&city, query, id)
	if err != nil {
		return City{}, err
	}
	return city, nil
}

func (s *SQLStorage) ListCities() ([]City, error) {
	var cities []City
	query := `SELECT * FROM cities`
	err := s.db.Select(&cities, query)
	if err != nil {
		return nil, err
	}
	return cities, nil
}
