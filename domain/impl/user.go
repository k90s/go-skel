package impl

import (
	"database/sql"

	"github.com/kai-ding/go-skel/domain"
	_ "github.com/lib/pq"
)

// UserService represents a PostgreSQL implementation of myapp.UserService.
type UserService struct {
	DB *sql.DB
}

func (s *UserService) User(id int) (*domain.User, error) {
	var u domain.User
	row := s.DB.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) Users() ([]*domain.User, error) {
	panic("not implemented")
}

func (s *UserService) CreateUser(u *domain.User) error {
	_, err := s.DB.Exec(`insert into users(name) values ($1)`, u.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	panic("not implemented")
}
