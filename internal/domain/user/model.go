package user

import "time"

// User representa a un usuario en el dominio.
type User struct {
	ID        string    // Identificador único
	Name      string    // Nombre completo
	Email     string    // Email único
	Password  string    // Password hashed
	CreatedAt time.Time // Fecha de creación
	UpdatedAt time.Time // Fecha de última actualización
}
