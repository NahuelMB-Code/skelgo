package user

import (
	"time"
)

type Service struct {
	repo UserRepository
}

func NewService(r UserRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateUser(u *User) (*User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}
