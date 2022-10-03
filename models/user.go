package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Fecha    string `json:"fecha"`
}
