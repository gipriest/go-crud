package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/XimenaGutierrezWagner/goCrud/models"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

type EmployeeServiceMock struct {
	mock.Mock
}

func (s *EmployeeServiceMock) Create(employee models.Employee) (models.EmployeeOutput, error) {
	args := s.Called(employee)

	t, ok := args.Get(0).(models.EmployeeOutput)
	if !ok {
		return models.EmployeeOutput{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeServiceMock) Update(employee *models.Employee, id int) (models.EmployeeOutput, error) {
	args := s.Called(employee, id)

	t, ok := args.Get(0).(models.EmployeeOutput)
	if !ok {
		return models.EmployeeOutput{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeServiceMock) GetAll() ([]models.Employee, error) {
	args := s.Called()

	t, ok := args.Get(0).([]models.Employee)
	if !ok {
		return []models.Employee{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeServiceMock) GetById(id int) (models.Employee, error) {
	args := s.Called(id)

	t, ok := args.Get(0).(models.Employee)
	if !ok {
		return models.Employee{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func (s *EmployeeServiceMock) Delete(id int) (models.EmployeeOutput, error) {
	args := s.Called(id)

	t, ok := args.Get(0).(models.EmployeeOutput)
	if !ok {
		return models.EmployeeOutput{}, errors.New("mock error")
	}

	return t, args.Error(1)
}

func NewEmployeeList() []models.Employee {
	return []models.Employee{}
}

func NewEmployee() models.Employee {
	return models.Employee{}
}

func TestEmployeeCreate_Success(t *testing.T) {
	output := models.EmployeeOutput{
		Completed: true,
	}

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Create", mock.AnythingOfType("Employee")).Return(output, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewBufferString(`{"first_name":"jhon","last_name": "Doe","date_of_birth": "1990-01-01","email": "test@gmail.com","phone_number": "3834112233","address": "Calle falsa 123","city": "SFVC","state": "Catamarca","country": "Argentina","postal_code": "4700","position": "test","department": "software","salary": 100.00}`))

	r := testRouter()
	r.POST("/employee", controller.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestEmployeeCreate_InvalidJSON(t *testing.T) {
	output := models.EmployeeOutput{
		Completed: true,
	}

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Create", mock.AnythingOfType("Employee")).Return(output, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewBufferString(`{"NOT_VALID_JSON"}`))

	r := testRouter()
	r.POST("/employee", controller.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestEmployeeCreate_failService(t *testing.T) {
	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Create", mock.AnythingOfType("Employee")).Return(models.EmployeeOutput{}, errors.New("unexpected error"))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewBufferString(`{"first_name":"jhon","last_name": "Doe","date_of_birth": "1990-01-01","email": "test@gmail.com","phone_number": "3834112233","address": "Calle falsa 123","city": "SFVC","state": "Catamarca","country": "Argentina","postal_code": "4700","position": "test","department": "software","salary": 100.00}`))

	r := testRouter()
	r.POST("/employee", controller.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestEmployeeUpdate_Success(t *testing.T) {
	id := 1

	output := models.EmployeeOutput{
		Completed: true,
	}

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Update", mock.AnythingOfType("*models.Employee"), id).Return(output, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/employee/1", bytes.NewBufferString(`{"first_name":"jhon","last_name": "Doe","date_of_birth": "1990-01-01","email": "test@gmail.com","phone_number": "3834112233","address": "Calle falsa 123","city": "SFVC","state": "Catamarca","country": "Argentina","postal_code": "4700","position": "test","department": "software","salary": 100.00}`))

	r := testRouter()
	r.PUT("/employee/:id", controller.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEmployeeUpdate_InvalidJSON(t *testing.T) {
	id := "1"
	output := models.EmployeeOutput{
		Completed: true,
	}

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Update", mock.AnythingOfType("*models.Employee"), id).Return(output, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/employee/1", bytes.NewBufferString(`{"NOT_VALID_JSON"}`))

	r := testRouter()
	r.PUT("/employee/:id", controller.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestEmployeeUpdate_failService(t *testing.T) {
	id := 1

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Update", mock.AnythingOfType("*models.Employee"), id).Return(models.EmployeeOutput{}, errors.New("unexpected error"))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/employee/1", bytes.NewBufferString(`{"first_name":"jhon","last_name": "Doe","date_of_birth": "1990-01-01","email": "test@gmail.com","phone_number": "3834112233","address": "Calle falsa 123","city": "SFVC","state": "Catamarca","country": "Argentina","postal_code": "4700","position": "test","department": "software","salary": 100.00}`))

	r := testRouter()
	r.PUT("/employee/:id", controller.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestEmployeeGetAll_Success(t *testing.T) {
	output := NewEmployeeList()

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("GetAll").Return(output, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/employee", nil)

	r := testRouter()
	r.GET("/employee", controller.GetAll)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEmployeeGetAll_failService(t *testing.T) {
	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("GetAll").Return([]models.Employee{}, errors.New("unexpected error"))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/employee", nil)

	r := testRouter()
	r.GET("/employee", controller.GetAll)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestEmployeeGetById_Success(t *testing.T) {
	output := NewEmployee()

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("GetById", 1).Return(output, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/employee/1", nil)

	r := testRouter()
	r.GET("/employee/:id", controller.GetById)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEmployeeGetById_failService(t *testing.T) {
	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("GetById", 1).Return(models.Employee{}, errors.New("unexpected error"))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/employee/1", nil)

	r := testRouter()
	r.GET("/employee/:id", controller.GetById)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}

func TestEmployeeDelete_Success(t *testing.T) {
	output := models.EmployeeOutput{
		Completed: true,
	}

	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Delete", 1).Return(output, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/employee/1", nil)

	r := testRouter()
	r.DELETE("/employee/:id", controller.Delete)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEmployeeDelete_failService(t *testing.T) {
	serviceMock := new(EmployeeServiceMock)
	controller := NewEmployeeController(serviceMock)

	serviceMock.On("Delete", 1).Return(models.Employee{}, errors.New("unexpected error"))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/employee/1", nil)

	r := testRouter()
	r.DELETE("/employee/:id", controller.Delete)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadGateway, w.Code)
}
