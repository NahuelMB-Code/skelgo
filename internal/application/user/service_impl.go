package user

import "time"

type service struct {
	repo UserRepository
}

func (s *service) CreateUser(u *User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return s.repo.Create(u)
}
