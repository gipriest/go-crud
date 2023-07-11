package services

import (
	"errors"
	"testing"

	"github.com/XimenaGutierrezWagner/goCrud/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepositoryMock struct{
	mock.Mock
}

func (s *EmployeeRepositoryMock) Create(employee models.Employee) (models.EmployeeOutput, error) {
	args := s.Called(employee)

	t, ok := args.Get(0).(models.EmployeeOutput)
	if !ok {
		return models.EmployeeOutput{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeRepositoryMock) Update(employee models.Employee, id int) (models.EmployeeOutput, error) {
	args := s.Called(employee, id)

	t, ok := args.Get(0).(models.EmployeeOutput)
	if !ok {
		return models.EmployeeOutput{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeRepositoryMock) GetAll() ([]models.Employee, error) {
	args := s.Called()

	t, ok := args.Get(0).([]models.Employee)
	if !ok {
		return []models.Employee{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeRepositoryMock) GetById(id int) (models.Employee, error) {
	args := s.Called(id)

	t, ok := args.Get(0).(models.Employee)
	if !ok {
		return models.Employee{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeRepositoryMock) Delete(id int) (models.EmployeeOutput, error) {
	args := s.Called(id)

	t, ok := args.Get(0).(models.EmployeeOutput)
	if !ok {
		return models.EmployeeOutput{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func EmployeeOutput() models.EmployeeOutput {
	return models.EmployeeOutput{
		Completed: true,
	}
}

func NewEmployeeList() []models.Employee {
	return []models.Employee{}
}

func NewEmployee() models.Employee {
	return models.Employee{}
}

func NewEmployeeId() int {
	return 1
}


func TestEmployeeCreate_Success(t *testing.T) {
	output := EmployeeOutput()
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("Create", mock.AnythingOfType("models.Employee")).Return(output, nil)
	res, _ := service.Create(NewEmployee())

	assert.Equal(t, res.Completed, output.Completed)
}

func TestEmployeeCreate_RepositoryFail(t *testing.T) {
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("Create", mock.AnythingOfType("models.Employee")).Return(models.EmployeeOutput{}, errors.New("unexpected error"))
	_, err := service.Create(NewEmployee())

	assert.NotNil(t, err)
}

func TestEmployeeUpdate_Success(t *testing.T) {
	id := NewEmployeeId()

	employee := models.Employee{
		FirstName: "test",
		LastName: "test",
		DateOfBirth: "1990-01-01",
		Email: "test@gmail.com",
		PhoneNumber: "test",
		Address: "test",
		City: "test",
		State: "test",
		Country: "test",
		PostalCode: "test",
		Position: "test",
		Department: "test",
		Salary: 100.00,
	}

	output := EmployeeOutput()
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("Update", mock.AnythingOfType("models.Employee"), id).Return(output, nil)
	res, _ := service.Update(&employee, id)

	assert.Equal(t, res.Completed, output.Completed)
}

func TestEmployeeUpdate_RepositoryFail(t *testing.T) {
	id := NewEmployeeId()

	employee := models.Employee{
		FirstName: "test",
		LastName: "test",
		DateOfBirth: "1990-01-01",
		Email: "test@gmail.com",
		PhoneNumber: "test",
		Address: "test",
		City: "test",
		State: "test",
		Country: "test",
		PostalCode: "test",
		Position: "test",
		Department: "test",
		Salary: 100.00,
	}

	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("Update", mock.AnythingOfType("models.Employee"), id).Return(models.EmployeeOutput{}, errors.New("unexpected error"))
	_, err := service.Update(&employee, id)

	assert.NotNil(t, err)
}

func TestEmployeeGetAll_Success(t *testing.T) {
	employees := NewEmployeeList()
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("GetAll").Return(employees, nil)
	res, _ := service.GetAll()

	assert.NotNil(t, res)
}

func TestEmployeeGetAll_RepositoryFail(t *testing.T) {
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("GetAll").Return([]models.Employee{}, errors.New("unexpected error"))
	_, err := service.GetAll()

	assert.NotNil(t, err)
}

func TestEmployeeGetById_Success(t *testing.T) {
	id := NewEmployeeId()
	employee := NewEmployee()
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("GetById", id).Return(employee, nil)
	res, _ := service.GetById(NewEmployeeId())

	assert.NotNil(t, res)
}

func TestEmployeeGetById_RepositoryFail(t *testing.T) {
	id := NewEmployeeId()
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("GetById", id).Return(models.Employee{}, errors.New("unexpected error"))
	_, err := service.GetById(NewEmployeeId())

	assert.NotNil(t, err)
}

func TestEmployeeDelete_Success(t *testing.T) {
	id := NewEmployeeId()
	output := EmployeeOutput()
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("Delete", id).Return(output, nil)
	res, _ := service.Delete(id)

	assert.NotNil(t, res)
	assert.Equal(t, res.Completed, output.Completed)
}

func TestEmployeeDelete_RepositoryFail(t *testing.T) {
	id := NewEmployeeId()
	repositoryMock := new(EmployeeRepositoryMock)
	service := NewDefaultEmployeeService(repositoryMock)

	repositoryMock.On("Delete", id).Return(models.EmployeeOutput{}, errors.New("unexpected error"))
	_, err := service.Delete(id)

	assert.NotNil(t, err)
}