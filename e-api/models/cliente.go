package models

import "time"

type Cliente struct {
	ID                uint   `gorm:"primaryKey;auto_increment;"`
	NombreUsuario     string `json:"Nombre_Usuario"`
	Contraseña        string
	Nombre            string
	Apellidos         string
	CorreoElectronico string `json:"Correo_Electronico"`
	Edad              uint
	Estatura          float32
	Peso              float32
	IMC               float32
	GEB               float32
	DeletedAt         *time.Time `sql:"index" json:"-"`
	CreatedAt         *time.Time `sql:"index" json:"-"`
	UpdatedAt         *time.Time `sql:"index" json:"-"`
}
