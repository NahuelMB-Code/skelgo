package user

// UserRepository es la interfaz que debe implementar cualquier repo (Mongo, SQL, etc).
type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}
