package models

type Persona struct {
	Nombre    string `json:"firts_name"`
	Apellido  string `json:"last_name"`
	Edad      int
	Direccion string
	Telefono  string
	Activo    bool
}
