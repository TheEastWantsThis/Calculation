package userservice

import (
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
	GetAllCalculationsForUser(id string) (User, error)
}

type uSer struct {
	repo UserRepository
}

func (s *uSer) GetAllCalculationsForUser(id string) (User, error) {
	return s.repo.GetAllCalculationsForUser(id)
}

func NewUserService(r UserRepository) *uSer {
	return &uSer{repo: r}
}

func (s *uSer) CreateUser(user User) (User, error) {
	us := User{
		ID:        uuid.NewString(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}
	if err := s.repo.CreateUser(us); err != nil {
		return User{}, err
	}
	return us, nil
}

func (s *uSer) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *uSer) GetUserByID(id string) (User, error) {
	return s.repo.GetUserByID(id)
}

func (s *uSer) UpdateUser(id string, user User) (User, error) {
	upus, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}
	now := time.Now()

	upus.Email = user.Email
	upus.Password = user.Password
	upus.UpdatedAt = &now

	if err := s.repo.UpdateUser(upus); err != nil {
		return User{}, err
	}
	return upus, nil

}
func (s *uSer) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
