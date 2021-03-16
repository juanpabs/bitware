package models

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql" // GORM driver
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB
var tables = []interface{}{
	&Cliente{},
}

func init() {
	var DB_USER string = os.Getenv("MARIADB_USER")
	var DB_PASS string = os.Getenv("MARIADB_PASS")
	var DB_NAME string = os.Getenv("MARIADB_NAME")
	var DB_HOST string = os.Getenv("MARIADB_HOST")
	var DB_PORT string = os.Getenv("MARIADB_PORT")
	var i int

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	for {
		var err error
		if i >= 3000 {
			panic("could not connect to " + source)
		}

		time.Sleep(3 * time.Second)
		db, err = gorm.Open("mysql", source)
		if err != nil {
			log.Info("Retrying connection...", err)
			i++
			continue
		}
		db.DB().SetMaxIdleConns(0)
		db.DB().SetConnMaxLifetime(time.Second * 14400)
		log.Info("Connected DB")
		break
	}
	Migrate()
}

func Migrate() {
	for _, t := range tables {
		db.AutoMigrate(t)
	}
}

func Create(cliente *Cliente) *gorm.DB {
	var oldcliente Cliente

	db.Where("nombre_usuario = ?", cliente.NombreUsuario).First(&oldcliente)
	if !reflect.ValueOf(oldcliente).IsZero() {
		return &gorm.DB{Error: errors.New("Nombre de Usuario existente")}
	}
	db.Where("correo_electronico = ?", cliente.CorreoElectronico).First(&oldcliente)
	if !reflect.ValueOf(oldcliente).IsZero() {
		return &gorm.DB{Error: errors.New("Correo existente")}
	} else {
		return db.Create(cliente)
	}
}

func AgregarCliente(cliente *Cliente) (*Cliente, error) {
	clienteDB := Create(cliente)
	return cliente, clienteDB.Error
}

func ModificarCliente(cliente *Cliente, id string) (*Cliente, error) {
	UintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return cliente, errors.New("Id no válido")
	}
	newClient := Cliente{}
	db.Where("id = ?", UintId).First(&newClient)
	if reflect.ValueOf(newClient).IsZero() {
		return cliente, errors.New("Usuario Inexistente")
	}
	newClient.setNewValues(cliente)
	clienteDB := db.Save(newClient)
	return &newClient, clienteDB.Error
}
func (c *Cliente) setNewValues(cliente *Cliente) {
	if cliente.Nombre != "" {
		c.Nombre = cliente.Nombre
	}
	if cliente.Apellidos != "" {
		c.Apellidos = cliente.Apellidos
	}
	if cliente.Contraseña != "" {
		c.Contraseña = cliente.Contraseña
	}
	if cliente.Edad != 0 {
		c.Edad = cliente.Edad
	}
	if cliente.Estatura != 0 {
		c.Estatura = cliente.Estatura
	}
	if cliente.GEB != 0 {
		c.GEB = cliente.GEB
	}
	if cliente.IMC != 0 {
		c.IMC = cliente.IMC
	}
	if cliente.Peso != 0 {
		c.Peso = cliente.Peso
	}
}

func GetTodosLosClientes() ([]Cliente, error) {
	clientes := []Cliente{}
	db.Find(&clientes)
	return clientes, db.Error
}

func GetClientePorId(id string) (Cliente, error) {
	cliente := Cliente{}
	UintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return cliente, errors.New("Id no válido")
	}

	db.Where("id = ?", UintId).First(&cliente)
	if reflect.ValueOf(cliente).IsZero() {
		return cliente, errors.New("Usuario Inexistente")
	}

	return cliente, db.Error
}
