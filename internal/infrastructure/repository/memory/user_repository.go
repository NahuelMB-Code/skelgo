package memory

import (
    "errors"
    "sync"

    "github.com/google/uuid"
    "github.com/NahuelMB-Code/skelgo/internal/application/user"
)

type UserRepository struct {
    mu    sync.Mutex
    store map[string]*user.User
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        store: make(map[string]*user.User),
    }
}

func (r *UserRepository) Create(u *user.User) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    // Generar un ID UUID para el usuario si no tiene
    if u.ID == "" {
        u.ID = uuid.NewString()
    }

    // Verificar si ya existe
    if _, exists := r.store[u.ID]; exists {
        return errors.New("user already exists")
    }

    // Guardar usuario
    r.store[u.ID] = u
    return nil
}

func (r *UserRepository) GetByID(id string) (*user.User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    u, ok := r.store[id]
    if !ok {
        return nil, errors.New("user not found")
    }
    return u, nil
}

func (r *UserRepository) Update(u *user.User) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.store[u.ID]; !exists {
        return errors.New("user not found")
    }

    r.store[u.ID] = u
    return nil
}

func (r *UserRepository) Delete(id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.store[id]; !exists {
        return errors.New("user not found")
    }

    delete(r.store, id)
    return nil
}
