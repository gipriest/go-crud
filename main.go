package main

import (
	"fmt"

	"github.com/XimenaGutierrezWagner/goCrud/controllers"
	"github.com/XimenaGutierrezWagner/goCrud/repository"
	"github.com/XimenaGutierrezWagner/goCrud/services"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func main() {

	// Defer database connection
	defer repository.DisconnectDB()

	// Database connect
	repository.InitDb()

	repository := repository.NewDefaultEmployeeRepository()
	serviceEmployee := services.NewDefaultEmployeeService(repository)
	controllerEmployee := controllers.NewEmployeeController(serviceEmployee)

	// Create HTTP router and start
	server = gin.Default()
	employee := server.Group("/employee")
	employee.POST("", controllerEmployee.Create)
	employee.PUT(":id", controllerEmployee.Update)
	employee.GET("", controllerEmployee.GetAll)
	employee.GET(":id", controllerEmployee.GetById)
	employee.DELETE(":id", controllerEmployee.Delete)
	fmt.Println(server.Run(":8080"))
}
