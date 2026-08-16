package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/NahuelMB-Code/skelgo/internal/application/user"
	"github.com/NahuelMB-Code/skelgo/internal/infrastructure/repository/memory"
	"github.com/NahuelMB-Code/skelgo/internal/interfaces/http/handler"
)

func main() {
	// Crear router Gorilla Mux
	r := mux.NewRouter()

	// Crear repositorio (en memoria para testear)
	userRepo := memory.NewUserRepository()

	// Crear caso de uso de usuario
	userUseCase := user.NewService(userRepo)

	// Crear handler HTTP de usuario
	userHandler := handler.NewUserHandler(userUseCase)

	// Registrar ruta POST /users
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	// Levantar servidor en puerto 8080
	fmt.Println("🚀 Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
