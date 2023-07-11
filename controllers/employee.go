package controllers

import (
	"net/http"
	"strconv"

	"github.com/XimenaGutierrezWagner/goCrud/models"
	"github.com/XimenaGutierrezWagner/goCrud/services"
	"github.com/gin-gonic/gin"
)

type Employee interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
}

type EmployeeController struct {
	EmployeeService services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) EmployeeController {
	return EmployeeController{
		EmployeeService: service,
	}
}

func (s EmployeeController) Create(c *gin.Context){
	var employee models.Employee
	
	err := c.ShouldBindJSON(&employee)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	output, err := s.EmployeeService.Create(employee)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (s EmployeeController) Update(c *gin.Context){
	var employee models.Employee

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"cannont convert id"})
		return
	}

	err = c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	output, err := s.EmployeeService.Update(&employee, id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (s EmployeeController) GetAll(c *gin.Context){
	employees, err := s.EmployeeService.GetAll()
	
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (s EmployeeController) GetById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"cannont convert id"})
		return
	}

	employee, err := s.EmployeeService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (s EmployeeController) Delete(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	output, err := s.EmployeeService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}