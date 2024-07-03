package service

import (
	"context"
	"github.com/Rochakb/go-stater-project/internal/model"
	"github.com/Rochakb/go-stater-project/internal/repository"
)

type EmployeeService interface {
	HealthCheck(ctx context.Context) string
	GetEmployee(ctx context.Context, empId int) (model.Employee, error)
	CreateEmployee(ctx context.Context, employee model.Employee) (bool, error)
	DeleteEmployee(ctx context.Context, empId int) (bool, error)
	UpdateEmployee(ctx context.Context, empId int, employee model.Employee) (bool, error)
	//GetEmployeesUnderBoss(ctx context.Context, bossId int) ([]model.Employee, error)
}

// Class Employee Service
type employeeService struct {
	repo repository.Repository
}

// Constructor of the class
func NewEmployeeService(repo repository.Repository) EmployeeService {
	return &employeeService{
		repo: repo,
	}
}

func (s *employeeService) HealthCheck(ctx context.Context) string {
	return "Hello, health check!"
}

func (s *employeeService) GetEmployee(ctx context.Context, empId int) (model.Employee, error) {
	employee, err := s.repo.GetEmployeeByID(ctx, empId)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (s *employeeService) CreateEmployee(ctx context.Context, employee model.Employee) (bool, error) {
	created, err := s.repo.CreateEmployee(ctx, employee)
	if err != nil {
		return false, err
	}
	return created, nil
}

func (s *employeeService) DeleteEmployee(ctx context.Context, empId int) (bool, error) {
	deleted, err := s.repo.DeleteEmployee(ctx, empId)
	if err != nil {
		return false, err
	}
	return deleted, nil
}

func (s *employeeService) UpdateEmployee(ctx context.Context, empId int, employee model.Employee) (bool, error) {
	updated, err := s.repo.UpdateEmployee(ctx, empId, employee)
	if err != nil {
		return false, err
	}
	return updated, nil
}

//func (s *employeeService) GetEmployeesUnderBoss(ctx context.Context, bossId int) ([]model.Employee, error) {
//	employees, err := s.repo.GetEmployeesUnderBoss(ctx, bossId)
//	if err != nil {
//		return nil, err
//	}
//	return employees, nil
//}
