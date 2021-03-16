package v1

import (
	"net/http"

	"github.com/bitware/e-api/models"
	"github.com/gin-gonic/gin"
)

func AgregarCliente(c *gin.Context) {
	request := models.Cliente{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	cliente, err := models.AgregarCliente(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, Respuesta(-1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, Respuesta(0, cliente))

}

func ModificarCliente(c *gin.Context) {
	request := models.Cliente{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")

	cliente, err := models.ModificarCliente(&request, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Respuesta(-1, err.Error()))
		return
	}

	c.JSON(http.StatusOK, Respuesta(0, cliente))

}

func GetCliente(c *gin.Context) {
	id := c.Param("id")

	if id == "todos" {
		clientes, err := models.GetTodosLosClientes()
		if err != nil {
			c.JSON(http.StatusNotFound, Respuesta(-1, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, Respuesta(0, clientes))
			return
		}
	} else {
		cliente, err := models.GetClientePorId(id)
		if err != nil {
			c.JSON(http.StatusNotFound, Respuesta(-1, err.Error()))
			return
		} else {
			c.JSON(http.StatusOK, Respuesta(0, cliente))
			return
		}
	}
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, "NUTRINET API REST")
}

func Respuesta(codigo int, mensaje interface{}) gin.H {
	return gin.H{
		"Cve_Error":   codigo,
		"Cve_Mensaje": mensaje,
	}
}
