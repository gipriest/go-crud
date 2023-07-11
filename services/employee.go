package services

import (
	"github.com/XimenaGutierrezWagner/goCrud/models"
	"github.com/XimenaGutierrezWagner/goCrud/repository"
)

type EmployeeService interface {
	Create(employee models.Employee) (models.EmployeeOutput, error)
	Update(employee *models.Employee, id int) (models.EmployeeOutput, error)
	GetAll() ([]models.Employee, error)
	GetById(id int) (models.Employee, error)
	Delete(id int) (models.EmployeeOutput, error)
}

type DefaultEmployeeService struct {
	repository repository.EmployeeRepository
}

func NewDefaultEmployeeService(repository repository.EmployeeRepository) DefaultEmployeeService {
	return DefaultEmployeeService{
		repository: repository,
	}
}

func (s DefaultEmployeeService) Create(employee models.Employee) (models.EmployeeOutput, error){
	res, err := s.repository.Create(employee)

	if err != nil {
		return models.EmployeeOutput{}, err
	}

	return res, nil
}

func (s DefaultEmployeeService) Update(employee *models.Employee, id int) (models.EmployeeOutput, error){
	res, err := s.repository.Update(*employee, id)
	
	if err != nil {
		return models.EmployeeOutput{}, err
	}

	return res, nil
}


func (s DefaultEmployeeService) GetAll()([]models.Employee, error){
	res, err := s.repository.GetAll()
	
	if err != nil {
		return []models.Employee{}, err
	}

	return res, nil
}

func (s DefaultEmployeeService) GetById(id int) (models.Employee, error){
	res, err := s.repository.GetById(id)
	
	if err != nil {
		return models.Employee{}, err
	}

	return res, nil
}

func (s DefaultEmployeeService) Delete(id int) (models.EmployeeOutput, error){
	res, err := s.repository.Delete(id)
	
	if err != nil {
		return models.EmployeeOutput{}, err
	}

	return res, nil
}