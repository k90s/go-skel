package impl

import (
	"database/sql"

	skel "github.com/kai-ding/go-skel"
	_ "github.com/lib/pq"
)

// UserService represents a PostgreSQL implementation of myapp.UserService.
type UserService struct {
	DB *sql.DB
}

func (s *UserService) User(id int) (*skel.User, error) {
	var u skel.User
	row := s.DB.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) Users() ([]*skel.User, error) {
	panic("not implemented")
}

func (s *UserService) CreateUser(u *skel.User) error {
	panic("not implemented")
}

func (s *UserService) DeleteUser(id int) error {
	panic("not implemented")
}
