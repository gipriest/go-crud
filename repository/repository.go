package repository

import (
	"errors"

	"github.com/XimenaGutierrezWagner/goCrud/models"
)

type EmployeeRepository interface {
	Create(employee models.Employee) (models.EmployeeOutput, error)
	Update(employee models.Employee, id int) (models.EmployeeOutput, error)
	GetAll() ([]models.Employee, error)
	GetById(id int) (models.Employee, error)
	Delete(id int) (models.EmployeeOutput, error)
}

type DefaultEmployeeRepository struct {
}

func NewDefaultEmployeeRepository() DefaultEmployeeRepository {
	return DefaultEmployeeRepository{}
}

func (r DefaultEmployeeRepository) Create(employee models.Employee) (models.EmployeeOutput, error) {
	err := Client.Create(&employee)
	if err.Error != nil {
		return models.EmployeeOutput{}, errors.New("unable to create")
	}

	output := models.EmployeeOutput{
		Completed: true,
	}

	return output, nil
}

func (r DefaultEmployeeRepository) Update(employee models.Employee, id int) (models.EmployeeOutput, error) {
	
	var employeeFind models.Employee

	findResult  := Client.First(&employeeFind, id)
	if findResult .Error != nil {
		return models.EmployeeOutput{}, errors.New("cannot find")
	}

	employeeFind.FirstName = employee.FirstName
	employeeFind.LastName = employee.LastName
	employeeFind.Address = employee.Address
	employeeFind.City = employee.City
	employeeFind.Country = employee.Country
	employeeFind.DateOfBirth = employee.DateOfBirth
	employeeFind.Department = employee.Department
	employeeFind.Email = employee.Email
	employeeFind.HireDate = employee.HireDate
	employeeFind.PhoneNumber = employee.PhoneNumber
	employeeFind.Position = employee.Position
	employeeFind.PostalCode = employee.PostalCode
	employeeFind.Salary = employee.Salary
	employeeFind.State = employee.State

	saveResult := Client.Save(&employeeFind)

	if saveResult.Error != nil {
		return models.EmployeeOutput{}, errors.New("cannot update")
	}

	output := models.EmployeeOutput{
		Completed: true,
	}

	return output, nil
}

func (r DefaultEmployeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	result := Client.Find(&employees)

	if result.Error != nil {
		return []models.Employee{}, errors.New("unable to find")
	}

	return employees, nil
}

func (r DefaultEmployeeRepository) GetById(id int) (models.Employee, error) {
	var employee models.Employee
	result := Client.First(&employee, id)

	if result.Error != nil {
		return models.Employee{}, errors.New("unable to find")
	}

	return employee, nil
}

func (r DefaultEmployeeRepository) Delete(id int) (models.EmployeeOutput, error) {
	var employee models.Employee
	result := Client.First(&employee, id)

	if result.Error != nil {
		return models.EmployeeOutput{}, errors.New("unable to find")
	}

	err := Client.Delete(&employee).Error

	if err != nil {
		return models.EmployeeOutput{}, err
	}

	output := models.EmployeeOutput{
		Completed: true,
	}

	return output, nil
}
