package v1

import "github.com/gin-gonic/gin"

var (
	Router *gin.Engine
	v1     *gin.RouterGroup
)

func init() {
	Router = gin.Default()
	Router.GET("/", Index)
	v1 = Router.Group("v1")
	NutriNet()
}

func Run(addr ...string) error {
	return Router.Run(addr...)
}

func NutriNet() {
	nutriNET := v1.Group("/NutriNET")
	{
		nutriNET.POST("/Cliente", AgregarCliente)
		nutriNET.GET("/Cliente/:id", GetCliente)
		nutriNET.PUT("/Cliente/:id", ModificarCliente)
	}

}
